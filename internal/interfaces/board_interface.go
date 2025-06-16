package interfaces

import "github.com/Gvardmeister/seabattle/internal/domain/coordinate"

type Board interface {
	PlaceShips()
	AllShipsKill() bool
	Render()
	IsAlreadyShot(x, y int) bool
	ReceiveShot(coord coordinate.Coordinate) bool
}
