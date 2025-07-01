package board_test

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/domain/ship"
)

func TestBoard_ReceiveShot_Miss(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	missCoord := coordinate.NewCoordinate(4, 5)

	hit := b.ReceiveShot(missCoord)

	if hit {
		t.Errorf("Ожидался промах, но вернуло true")
	}
	if b.Grid[4][5] != board.Loss {
		t.Errorf("Ожидалось, что клетка станет Loss, но стало %v", b.Grid[4][5])
	}
}

func TestBoard_ReceiveShot_AlreadyShot(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	coord := coordinate.NewCoordinate(1, 1)

	b.Grid[1][1] = board.Loss

	hit := b.ReceiveShot(coord)

	if hit {
		t.Errorf("Ожидалось false, потому что сюда уже стреляли")
	}
}

func TestBoard_ReceiveShot_Hit(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	shipCoord := coordinate.NewCoordinate(2, 3)

	b.Ships = []ship.Ship{
		{
			Cells: []coordinate.Coordinate{shipCoord},
			Hit:   []bool{false},
		},
	}

	hit := b.ReceiveShot(shipCoord)

	if !hit {
		t.Errorf("Ожидалось попадание, но вернуло false")
	}

	if b.Grid[2][3] != board.Deck {
		t.Errorf("Ожидалось, что клетка станет Deck, но стало %v", b.Grid[2][3])
	}

	if !b.Ships[0].Hit[0] {
		t.Errorf("Ожидалось, что Hit будет true, но он false")
	}
}
