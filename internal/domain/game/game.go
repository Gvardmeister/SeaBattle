package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/domain/player"
)

type game struct {
	player1 player.Player
	board1  board.Board
	player2 player.Player
	board2  board.Board
}

func NewGame() *game {
	fmt.Println("Игра морской бой!")

	reader := bufio.NewReader(os.Stdin)
	name1 := getPlayerName(reader, "Введите имя первого игрока: ")
	name2 := getPlayerName(reader, "Введите имя второго игрока: ")

	return &game{
		player1: player.NewHumanPlayer(name1),
		board1:  *board.NewBoard(board.SizeBoard),
		player2: player.NewHumanPlayer(name2),
		board2:  *board.NewBoard(board.SizeBoard),
	}
}

func getPlayerName(reader *bufio.Reader, name string) string {
	fmt.Print(name)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (g *game) MakeMove(coord coordinate.Coordinate, enemyBoard *board.Board) {
	x, y := coord.GetX(), coord.GetY()
	hit := false

	for i := range enemyBoard.Ships {
		ship := &enemyBoard.Ships[i]

		for j, cell := range ship.Cells {
			if cell.GetX() == x && cell.GetY() == y {
				ship.Hit[j] = true
				hit = true

				break
			}
		}

		if hit {
			break
		}
	}

	if hit {
		enemyBoard.Grid[x][y] = board.Deck // переставил местами для удобства пользователя

		fmt.Println("\nПопадание!")
	} else {
		enemyBoard.Grid[x][y] = board.Loss // переставил местами для удобства пользователя

		fmt.Println("\nВы промахнулись!")
	}

	fmt.Println()
	enemyBoard.Render()
}

func (g *game) turn(p player.Player, enemyBoard *board.Board) bool {
	fmt.Printf("\nХод игрока: %s\n", p.GetName())
	enemyBoard.Render()

	coord, wantsExit := p.GetMove(enemyBoard)

	if wantsExit {
		fmt.Printf("\nИгрок %s сдался. Игра окончена!\n", p.GetName())
		return true
	}

	g.MakeMove(coord, enemyBoard)
	return false
}

func (g *game) Start() {
	g.board1.PlaceShips()
	g.board2.PlaceShips()

	for {
		if g.turn(g.player1, &g.board2) {
			break
		}

		if g.board2.AllShipsKill() {
			fmt.Printf("\nПобедил игрок - %s", g.player1.GetName())
			break
		}

		if g.turn(g.player2, &g.board1) {
			break
		}

		if g.board1.AllShipsKill() {
			fmt.Printf("\nПобедил игрок - %s", g.player2.GetName())
			break
		}
	}
}
