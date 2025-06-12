package board_test

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/domain/ship"
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

func TestAllShipsKillOne(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	cells := []coordinate.Coordinate{
		coordinate.NewCoordinate(0, 0),
	}

	s := ship.NewShip(cells)
	b.Ships = append(b.Ships, *s)

	if b.AllShipsKill() {
		t.Errorf("Ожидалось false, получено true, так как корабль не подбит")
	}
}

func TestAllShipsKillNoShips(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)

	if !b.AllShipsKill() {
		t.Errorf("Ожидалось true, получено false - на поле нет кораблей")
	}
}

func TestAllShipsKillOneDead(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	cells := []coordinate.Coordinate{
		coordinate.NewCoordinate(0, 0),
		coordinate.NewCoordinate(0, 1),
	}

	s := ship.NewShip(cells)
	for i := range s.Hit {
		s.Hit[i] = true
	}

	b.Ships = append(b.Ships, *s)

	if !b.AllShipsKill() {
		t.Errorf("Ожидалось true, получено false - один корабль полностью подбит")
	}
}

func TestAllShipsKillOnePartDead(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	cells := []coordinate.Coordinate{
		coordinate.NewCoordinate(0, 0),
		coordinate.NewCoordinate(0, 1),
	}

	s := ship.NewShip(cells)
	s.Hit[0] = true

	b.Ships = append(b.Ships, *s)

	if b.AllShipsKill() {
		t.Errorf("Ожидалось false, получено true - корабль подбит частично")
	}
}

func TestAllShipsKillOneDeadTwoLive(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	cells1 := []coordinate.Coordinate{
		coordinate.NewCoordinate(0, 0),
		coordinate.NewCoordinate(0, 1),
	}
	s1 := ship.NewShip(cells1)
	s1.Hit[0] = true
	s1.Hit[1] = true

	cells2 := []coordinate.Coordinate{
		coordinate.NewCoordinate(2, 0),
		coordinate.NewCoordinate(2, 1),
	}
	s2 := ship.NewShip(cells2)

	b.Ships = append(b.Ships, *s1, *s2)

	if b.AllShipsKill() {
		t.Errorf("Ожидалось false, получено true - второй корабль не подбит")
	}
}

func TestAllShipsKillTwoDeadOneLive(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	cells1 := []coordinate.Coordinate{
		coordinate.NewCoordinate(0, 0),
		coordinate.NewCoordinate(0, 1),
	}
	s1 := ship.NewShip(cells1)

	cells2 := []coordinate.Coordinate{
		coordinate.NewCoordinate(2, 0),
		coordinate.NewCoordinate(2, 1),
	}
	s2 := ship.NewShip(cells2)
	s2.Hit[0] = true
	s2.Hit[1] = true

	b.Ships = append(b.Ships, *s1, *s2)

	if b.AllShipsKill() {
		t.Errorf("Ожидалось false, получено true - первый корабль не подбит")
	}
}

func TestAllShipsKillDead(t *testing.T) {
	b := board.NewBoard(board.SizeBoard)
	cells1 := []coordinate.Coordinate{
		coordinate.NewCoordinate(0, 0),
		coordinate.NewCoordinate(0, 1),
	}
	s1 := ship.NewShip(cells1)
	s1.Hit[0] = true
	s1.Hit[1] = true

	cells2 := []coordinate.Coordinate{
		coordinate.NewCoordinate(2, 0),
		coordinate.NewCoordinate(2, 1),
	}
	s2 := ship.NewShip(cells2)
	s2.Hit[0] = true
	s2.Hit[1] = true

	b.Ships = append(b.Ships, *s1, *s2)

	if !b.AllShipsKill() {
		t.Errorf("Ожидалось true, получено false - оба корабля подбиты")
	}
}

func TestIsOutOfBounds(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected bool
	}{
		{"X < 0", -1, 0, true},
		{"Y < 0", 0, -1, true},
		{"X == SizeBoard", board.SizeBoard, 0, true},
		{"Y == SizeBoard", 0, board.SizeBoard, true},
		{"X and Y < 0", -1, -1, true},
		{"X and Y == SizeBoard", board.SizeBoard, board.SizeBoard, true},
		{"Валидные (0,0)", 0, 0, false},
		{"Валидные (SizeBoard-1,SizeBoard-1)", board.SizeBoard - 1, board.SizeBoard - 1, false},
		{"Валидные (5,5)", 5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := board.IsOutOfBounds(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("IsOutOfBounds(%d, %d) = %v; ожидалось %v", tt.x, tt.y, result, tt.expected)
			}
		})
	}
}
