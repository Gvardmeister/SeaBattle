package player

import "github.com/Gvardmeister/seabattle/internal/domain/coordinate"

type Player interface {
	GetMove(grid [][]string) coordinate.Coordinate
}
