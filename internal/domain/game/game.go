package game

import (
	"fmt"

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
	fmt.Println("Чтобы сдаться/продолжить игру введите - y/n")

	var name1, name2 string

	fmt.Print("Введите имя первого игрока: ")
	fmt.Scan(&name1)
	fmt.Print("Введите имя второго игрока: ")
	fmt.Scan(&name2)

	return &game{
		player1: player.NewHumanPlayer(name1),
		board1:  *board.NewBoard(board.SizeBoard),
		player2: player.NewHumanPlayer(name2),
		board2:  *board.NewBoard(board.SizeBoard),
	}
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

	if p.StopGame() {
		fmt.Printf("\n%s, завершил игру!", p.GetName())
		return true
	}

	coord := p.GetMove(enemyBoard)
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
