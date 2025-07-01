package interfaces

import (
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type Player interface {
	GetMove(b Board) (coordinate.Coordinate, bool)
	GetName() string
}
