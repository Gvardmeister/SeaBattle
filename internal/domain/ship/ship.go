package ship

import "github.com/Gvardmeister/seabattle/internal/domain/coordinate"

type Ship struct {
	Cells []coordinate.Coordinate
	Hit   []bool
}

func NewShip(cells []coordinate.Coordinate) *Ship {
	return &Ship{
		Cells: cells,
		Hit:   make([]bool, len(cells)),
	}
}
