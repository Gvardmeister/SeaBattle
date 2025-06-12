package board_test

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
)

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
