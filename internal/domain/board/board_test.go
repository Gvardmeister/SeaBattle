package board_test

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
)

func TestIsAlreadyShot(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	b.Grid[0][0] = "@"
	b.Grid[0][1] = "."
	b.Grid[0][2] = "X"

	if b.IsAlreadyShot(0, 0) {
		t.Errorf("Ожидалось false в (0, 0), получено true")
	}

	if !b.IsAlreadyShot(0, 1) {
		t.Errorf("Ожидалось true в (0, 1), получено false")
	}

	if !b.IsAlreadyShot(0, 2) {
		t.Errorf("Ожидалось true в (0, 2), полученого false")
	}
}
