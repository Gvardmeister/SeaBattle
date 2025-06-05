package ship

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

func TestNewShip(t *testing.T) {
	cells := []coordinate.Coordinate{
		coordinate.NewCoordinate(1, 2),
		coordinate.NewCoordinate(2, 1),
		coordinate.NewCoordinate(1, 1),
		coordinate.NewCoordinate(-1, 1),
		coordinate.NewCoordinate(1, -1),
	}

	ship := NewShip(cells)

	if len(cells) != len(ship.Cells) {
		t.Errorf("Ожидалась длина Cells = %d, получено %d", len(cells), len(ship.Cells))
	}
	if len(cells) != len(ship.Hit) {
		t.Errorf("Ожидалась длина Hit = %d, получено %d", len(cells), len(ship.Hit))
	}

	for i, hit := range ship.Hit {
		if hit {
			t.Errorf("Hit[%d] должен быть false по умолчанию", i)
		}
	}
}
