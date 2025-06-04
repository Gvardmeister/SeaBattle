package player

import (
	"fmt"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/pkg/interfaces"
)

type BotPlayer struct {
	name     string
	coordGen interfaces.CoordinateGenerator
}

func NewBotPlayer(name string, gen interfaces.CoordinateGenerator) *BotPlayer {
	return &BotPlayer{
		name:     name,
		coordGen: gen,
	}
}

func (bp *BotPlayer) GetName() string {
	return bp.name
}

func (bp *BotPlayer) GetMove(b *board.Board) (coordinate.Coordinate, bool) {
	coord := bp.coordGen.Generator(b)
	fmt.Printf("\nХод бота: %s %d %d\n", bp.name, coord.GetX()+1, coord.GetY()+1)
	return coord, false
}
