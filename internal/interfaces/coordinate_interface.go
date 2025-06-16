package interfaces

import (
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type CoordinateGenerator interface {
	Generator(b Board) coordinate.Coordinate
}
