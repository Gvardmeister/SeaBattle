package game

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockplayer"
)

func TestGame_turn_PlayerWantsExit(t *testing.T) {
	mockPlayer := &mockplayer.MockPlayer{
		Name: "playerExit",
		GetMoveFunc: func(b *board.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(0, 0), true
		},
		GetNameFunc: func() string {
			return "playerExit"
		},
	}

	g := &game{}
	board := board.NewBoard(board.SizeBoard)

	turn := g.turn(mockPlayer, board)

	if !turn {
		t.Errorf("Ожидалось true, так как игрок сдался, но получили %v", turn)
	}
}

func TestGame_turn_PlayerMakesMove(t *testing.T) {
	mockPlayer := &mockplayer.MockPlayer{
		Name: "playerMove",
		GetMoveFunc: func(b *board.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(1, 1), false
		},
		GetNameFunc: func() string { return "playerMove" },
	}

	g := &game{}
	board := board.NewBoard(board.SizeBoard)

	turn := g.turn(mockPlayer, board)

	if turn {
		t.Errorf("Ожидалось false, так как игрок сделал ход, но получили %v", turn)
	}
}
