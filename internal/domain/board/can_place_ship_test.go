package board

import (
	"testing"
)

func TestCanPlaceShip_ValidHorizontalPlacement(t *testing.T) {
	b := NewBoard(SizeBoard)

	x, y := 2, 2
	size := 3
	orientation := Horizontal

	if !b.canPlaceShip(x, y, orientation, size) {
		t.Errorf("Ожидалось размещение корабля в (%d, %d), получено", x, y)
	}
}

func TestCanPlaceShip_ValidVerticalPlacement(t *testing.T) {
	b := NewBoard(SizeBoard)

	x, y := 2, 2
	size := 3
	orientation := Vertical

	if !b.canPlaceShip(x, y, orientation, size) {
		t.Errorf("Ожидалось размещение корабля вертикально в (%d, %d), но функция вернула false", x, y)
	}
}

func TestCanPlaceShip_ValidRangeX(t *testing.T) {
	b := NewBoard(SizeBoard)

	x, y := 8, 0
	size := 4
	orientation := Vertical

	if b.canPlaceShip(x, y, orientation, size) {
		t.Errorf("Ожидалось, что нельзя будет поставить корабль за границей по оси Х")
	}
}

func TestCanPlaceShip_ValidRangeY(t *testing.T) {
	b := NewBoard(SizeBoard)

	x, y := 0, 8
	size := 4
	orientation := Horizontal

	if b.canPlaceShip(x, y, orientation, size) {
		t.Errorf("Ожидалось, что нельзя будет поставить корабль за границей по оси Y")
	}
}

func TestCanPlaceShip_Merge(t *testing.T) {
	b := NewBoard(SizeBoard)

	b.placeShip(0, 0, Horizontal, 4)

	if b.canPlaceShip(0, 0, Vertical, 4) {
		t.Errorf("Ожидалось, что корабль нельзя ставить в эту позицию, т.к. он соприкасается с другим кораблём")
	}
}

func TestCanPlaceShip_TouchingSide(t *testing.T) {
	b := NewBoard(SizeBoard)

	b.placeShip(0, 0, Horizontal, 4)

	if b.canPlaceShip(1, 1, Horizontal, 2) {
		t.Errorf("Ожидалось, что корабль не может быть размещён вплотную к другому")
	}
}

func TestCanPlaceShip_TouchingDiagonal(t *testing.T) {
	b := NewBoard(SizeBoard)

	b.placeShip(1, 1, Horizontal, 3)

	if b.canPlaceShip(0, 0, Horizontal, 2) {
		t.Errorf("Ожидалось, что корабль не может быть размещён по диагонали к другому")
	}
}
