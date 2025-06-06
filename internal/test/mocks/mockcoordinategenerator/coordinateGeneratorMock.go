package mocks

import (
	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type MockCoordinateGenerator struct {
	FixedCoord coordinate.Coordinate
	Called     bool
}

func (m *MockCoordinateGenerator) Generator(b *board.Board) coordinate.Coordinate {
	m.Called = true
	return m.FixedCoord
}
