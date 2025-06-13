package board

import (
	"testing"
)

func TestCanPlaceShip_ValidHorizontalPlacement(t *testing.T) {
	b := NewBoard(SizeBoard)

	x, y := 2, 2
	size := 3
	orientation := Horizontal

	if !b.сanPlaceShip(x, y, orientation, size) {
		t.Errorf("Ожидалось размещение корабля в (%d, %d), получено", x, y)
	}
}

func TestCanPlaceShip_ValidVerticalPlacement(t *testing.T) {
	b := NewBoard(SizeBoard)

	x, y := 2, 2
	size := 3
	orientation := Vertical

	if !b.сanPlaceShip(x, y, orientation, size) {
		t.Errorf("Ожидалось размещение корабля в (%d, %d), получено", x, y)
	}
}
