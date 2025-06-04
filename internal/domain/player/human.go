package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type HumanPlayer struct {
	name string
}

func NewHumanPlayer(name string) *HumanPlayer {
	return &HumanPlayer{name: name}
}

func (hp *HumanPlayer) GetMove(b *board.Board) (coordinate.Coordinate, bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВведите координаты. Формат: {1 1} или 'y' для выхода:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "y" {
			return coordinate.NewCoordinate(-1, -1), true
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("\nНекорректный ввод. Введите, например: 3 5 или 'y' для выхода.")
			continue
		}

		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("\nВведите два целых числа или 'y' для выхода.")
			continue
		}

		if x < 1 || x > board.SizeBoard || y < 1 || y > board.SizeBoard {
			fmt.Println("\nКоординаты вне диапазона. Повторите ввод.")
			continue
		}

		x--
		y--

		if !b.IsAlreadyShot(x, y) {
			return coordinate.NewCoordinate(x, y), false
		}

		fmt.Println("\nВы уже стреляли в эту клетку.")
	}
}

func (hp *HumanPlayer) GetName() string {
	return hp.name
}
