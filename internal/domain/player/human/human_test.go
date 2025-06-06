package human

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockinputreader"
)

func TestNewHumanPlayer(t *testing.T) {
	inputs := []string{"3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)

	if player.GetName() != "Alex" {
		t.Errorf("Ожидалось имя 'Alex', получено %s", player.GetName())
	}
}

func TestGetMove(t *testing.T) {
	inputs := []string{"3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}

	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d, %d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveExit(t *testing.T) {
	inputs := []string{"y"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if !exit {
		t.Errorf("Ожидалось exit == true, получено false")
	}

	if coord.GetX() != -1 || coord.GetY() != -1 {
		t.Errorf("Ожидались координаты (-1, -1), получены (%d, %d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveInvalidInput(t *testing.T) {
	inputs := []string{"asdsad", "3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}

	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d, %d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveEmptyInput(t *testing.T) {
	inputs := []string{"   ", "3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}

	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d, %d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveInvalidDuplicate(t *testing.T) {
	inputs := []string{"1 1", "2 2"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	b := board.NewBoard(board.SizeBoard)

	b.Grid[0][0] = board.Deck

	coord, exit := player.GetMove(b)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}

	if coord.GetX() != 1 || coord.GetY() != 1 {
		t.Errorf("Ожидались координаты (1, 1), получены (%d, %d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveArray(t *testing.T) {
	inputs := []string{"100 100", "3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}

	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d, %d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveSingleNumber(t *testing.T) {
	inputs := []string{"3", "3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}
	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d,%d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMovePartialInvalid(t *testing.T) {
	inputs := []string{"5 abc", "3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}
	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d,%d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveFloatInvalid(t *testing.T) {
	inputs := []string{"2.0 3.0", "3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}
	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d,%d)", coord.GetX(), coord.GetY())
	}
}

func TestGetMoveRangeInvalid(t *testing.T) {
	inputs := []string{"3 4 6 7", "3 4"}
	mockReader := mockinputreader.NewMockInput(inputs)

	player := NewHumanPlayer("Alex", mockReader)
	board := board.NewBoard(board.SizeBoard)

	coord, exit := player.GetMove(board)
	if exit {
		t.Errorf("Ожидалось exit == false, получено true")
	}
	if coord.GetX() != 2 || coord.GetY() != 3 {
		t.Errorf("Ожидались координаты (2, 3), получены (%d,%d)", coord.GetX(), coord.GetY())
	}
}
