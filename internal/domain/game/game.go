package game

import (
	"fmt"
	"strings"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/domain/player/bot"
	"github.com/Gvardmeister/seabattle/internal/domain/player/human"
	"github.com/Gvardmeister/seabattle/internal/interfaces"
	"github.com/Gvardmeister/seabattle/internal/services/stdgenerator"
)

type game struct {
	player1 interfaces.Player
	board1  interfaces.Board
	player2 interfaces.Player
	board2  interfaces.Board
}

func NewGame(input interfaces.InputReader) *game {
	fmt.Println("Игра морской бой!")
	fmt.Println("Выберите режим: PvP - 1 или PvE - 2")

	modeInput, err := input.ReadLine()
	if err != nil {
		fmt.Println("Ошибка чтения. Выбран режим по умолчанию: PvE")
		modeInput = "2"
	}
	modeInput = strings.TrimSpace(modeInput)

	var player1, player2 interfaces.Player

	name1 := getPlayerName(input, "Введите имя первого игрока: ")
	player1 = human.NewHumanPlayer(name1, input)

	if modeInput == "1" {
		name2 := getPlayerName(input, "Введите имя второго игрока: ")
		player2 = human.NewHumanPlayer(name2, input)
	} else {
		player2 = bot.NewBotPlayer("Bot", getDefaultGenerator())
	}

	return &game{
		player1: player1,
		board1:  board.NewBoard(board.SizeBoard),
		player2: player2,
		board2:  board.NewBoard(board.SizeBoard),
	}
}

func (g *game) GetPlayer1() interfaces.Player {
	return g.player1
}

func (g *game) GetPlayer2() interfaces.Player {
	return g.player2
}

func getDefaultGenerator() interfaces.CoordinateGenerator {
	return stdgenerator.NewRandomGenerator()
}

func getPlayerName(reader interfaces.InputReader, name string) string {
	fmt.Print(name)
	input, err := reader.ReadLine()
	if err != nil || strings.TrimSpace(input) == "" {
		fmt.Println("Ошибка чтения. Используется имя по умолчанию")

		if strings.Contains(name, "второго") {
			return "Player2"
		}
		return "Player1"
	}
	return strings.TrimSpace(input)
}

func (g *game) MakeMove(coord coordinate.Coordinate, enemyBoard interfaces.Board) {
	hit := enemyBoard.ReceiveShot(coord)

	if hit {
		fmt.Println("\nПопадание!")
	} else {
		fmt.Println("\nВы промахнулись!")
	}

	fmt.Println()
	enemyBoard.Render()
}

func (g *game) turn(p interfaces.Player, enemyBoard interfaces.Board) bool {
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
		if g.turn(g.player1, g.board2) {
			break
		}

		if g.board2.AllShipsKill() {
			fmt.Printf("\nПобедил игрок - %s", g.player1.GetName())
			break
		}

		if g.turn(g.player2, g.board1) {
			break
		}

		if g.board1.AllShipsKill() {
			fmt.Printf("\nПобедил игрок - %s", g.player2.GetName())
			break
		}
	}
}
