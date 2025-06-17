package mockplayer

import (
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/interfaces"
)

type MockPlayer struct {
	Name        string
	Moved       bool
	GetNameFunc func() string
	GetMoveFunc func(b interfaces.Board) (coordinate.Coordinate, bool)
	TurnFunc    func(b interfaces.Board) bool
}

func (m *MockPlayer) GetName() string {
	if m.GetNameFunc != nil {
		return m.GetNameFunc()
	}
	return m.Name
}

func (m *MockPlayer) GetMove(b interfaces.Board) (coordinate.Coordinate, bool) {
	m.Moved = true
	if m.GetMoveFunc != nil {
		return m.GetMoveFunc(b)
	}
	return coordinate.Coordinate{}, false
}
