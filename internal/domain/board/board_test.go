package board_test

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
)

func TestIsAlreadyShot(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	b.Grid[0][0] = "@"
	b.Grid[0][1] = "."
	b.Grid[0][2] = "X"

	if b.IsAlreadyShot(0, 0) {
		t.Errorf("Ожидалось false в (0, 0), получено true")
	}

	if !b.IsAlreadyShot(0, 1) {
		t.Errorf("Ожидалось true в (0, 1), получено false")
	}

	if !b.IsAlreadyShot(0, 2) {
		t.Errorf("Ожидалось true в (0, 2), полученого false")
	}
}

func TestPlaceShips_CountAndValidity(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)

	b.PlaceShips()

	expectedShipsCount := 10
	if len(b.Ships) != expectedShipsCount {
		t.Fatalf("Ожидалось %d кораблей, получено %d", expectedShipsCount, len(b.Ships))
	}

	for _, ship := range b.Ships {
		for _, c := range ship.Cells {
			if c.GetX() < 0 || c.GetX() >= board.SizeBoard || c.GetY() < 0 || c.GetY() >= board.SizeBoard {
				t.Errorf("Корабль выходит за границы поля: %+v", ship)
			}
		}
	}

	for i, s1 := range b.Ships {
		for j, s2 := range b.Ships {
			if i == j {
				continue
			}
			for _, c1 := range s1.Cells {
				for _, c2 := range s2.Cells {
					dx := c1.GetX() - c2.GetX()
					dy := c1.GetY() - c2.GetY()
					if abs(dx) <= 1 && abs(dy) <= 1 {
						t.Errorf("Корабли соседствуют: индекс %d и %d", i, j)
					}
				}
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
