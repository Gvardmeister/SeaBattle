package player

import (
	"fmt"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type HumanPlayer struct{}

func (hp *HumanPlayer) GetMove(grid [][]string) coordinate.Coordinate {
	for {
		var x, y int
		fmt.Println("\nВведите координаты X и Y {от 1 до 10}:")
		fmt.Scan(&x, &y)

		if x < 1 || x > board.SizeBoard || y < 1 || y > board.SizeBoard {
			fmt.Println("\nНеверные координаты. Повторите ввод.")
			continue
		}

		x--
		y--

		if grid[x][y] == board.Change { // переставил местами для удобства пользователя
			return coordinate.NewCoordinate(x, y)
		}
		fmt.Println("\nВы уже стреляли в эту клетку.")
	}
}
