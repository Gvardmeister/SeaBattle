package mockcoordinategenerator

import (
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/interfaces"
)

type MockCoordinateGenerator struct {
	FixedCoord coordinate.Coordinate
	Called     bool
}

func (m *MockCoordinateGenerator) Generator(b interfaces.Board) coordinate.Coordinate {
	m.Called = true
	return m.FixedCoord
}
