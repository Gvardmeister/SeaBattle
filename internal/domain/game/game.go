package game

import (
	"fmt"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/domain/player"
)

type game struct {
	player player.Player
	board  board.Board
}

func NewGame() *game {
	return &game{
		player: &player.HumanPlayer{},
		board:  *board.NewBoard(board.SizeBoard),
	}
}

func (g *game) MakeMove(coord coordinate.Coordinate) {
	x, y := coord.GetX(), coord.GetY()
	hit := false

	for i := range g.board.Ships {
		ship := &g.board.Ships[i]

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
		g.board.Grid[x][y] = board.Deck // переставил местами для удобства пользователя

		fmt.Println("\nПопадание!")
	} else {
		g.board.Grid[x][y] = board.Loss // переставил местами для удобства пользователя

		fmt.Println("\nВы промахнулись!")
	}

	fmt.Println()
	g.board.Render()
}

func (g *game) InitGame() {
	g.board.PlaceShips()
	g.board.Render()

	for {
		coord := g.player.GetMove(&g.board)
		g.MakeMove(coord)

		if g.board.AllShipsKill() {
			fmt.Println("\nВсе корабли уничтожены!")
			break
		}
	}
}
