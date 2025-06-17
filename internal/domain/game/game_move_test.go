package game

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockboard"
)

func TestMakeMove_Hit(t *testing.T) {
	board := &mockboard.MockBoard{
		ReceiveShotFunc: func(c coordinate.Coordinate) bool {
			return true
		},
	}

	game := &game{}
	coord := coordinate.NewCoordinate(1, 1)

	game.MakeMove(coord, board)
}
