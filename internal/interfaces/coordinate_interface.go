package interfaces

import (
	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type CoordinateGenerator interface {
	Generator(b *board.Board) coordinate.Coordinate
}
