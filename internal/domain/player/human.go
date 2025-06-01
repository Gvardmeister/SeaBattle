package player

import (
	"fmt"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type HumanPlayer struct {
	name string
}

func NewHumanPlayer(name string) *HumanPlayer {
	return &HumanPlayer{
		name: name,
	}
}

func (hp *HumanPlayer) GetMove(b *board.Board) coordinate.Coordinate {
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

		if !b.IsAlreadyShot(x, y) { // переставил местами для удобства пользователя
			return coordinate.NewCoordinate(x, y)
		}

		fmt.Println("\nВы уже стреляли в эту клетку.")
	}
}

func (hp *HumanPlayer) GetName() string {
	return hp.name
}

func (hp *HumanPlayer) StopGame() bool {
	var quit string
	// fmt.Println("\nХочешь выйти из игры? (y/n):")
	fmt.Scan(&quit)

	return quit == "y"
}
