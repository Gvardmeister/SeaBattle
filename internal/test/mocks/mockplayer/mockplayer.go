package mockplayer

import (
	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type MockPlayer struct {
	Name        string
	GetNameFunc func() string
	GetMoveFunc func(*board.Board) (coordinate.Coordinate, bool)
}

func (m *MockPlayer) GetName() string {
	if m.GetNameFunc != nil {
		return m.GetNameFunc()
	}
	return m.Name
}

func (m *MockPlayer) GetMove(b *board.Board) (coordinate.Coordinate, bool) {
	if m.GetMoveFunc != nil {
		return m.GetMoveFunc(b)
	}
	return coordinate.Coordinate{}, false
}
