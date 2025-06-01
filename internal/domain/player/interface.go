package player

import (
	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type Player interface {
	GetMove(b *board.Board) (coordinate.Coordinate, bool)
	GetName() string
}
