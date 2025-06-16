package mockboard

import (
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type MockBoard struct {
	PlaceShipsCalled  bool
	AllShipsKillFunc  func() bool
	IsAlreadyShotFunc func(x, y int) bool
}

func (m *MockBoard) PlaceShips() {
	m.PlaceShipsCalled = true
}

func (m *MockBoard) AllShipsKill() bool {
	if m.AllShipsKillFunc != nil {
		return m.AllShipsKillFunc()
	}
	return false
}

func (m *MockBoard) IsAlreadyShot(x, y int) bool {
	if m.IsAlreadyShotFunc != nil {
		return m.IsAlreadyShotFunc(x, y)
	}
	return false
}

func (m *MockBoard) Render() {}

func (m *MockBoard) ReceiveShot(c coordinate.Coordinate) bool {
	return false
}
