package stdgenerator

import (
	"math/rand"
	"time"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
)

type RandomGenerator struct {
	rng *rand.Rand
}

func NewRandomGenerator() *RandomGenerator {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &RandomGenerator{
		rng: generator,
	}
}

func (rg *RandomGenerator) Generator(b *board.Board) coordinate.Coordinate {
	for {
		x := rg.rng.Intn(board.SizeBoard)
		y := rg.rng.Intn(board.SizeBoard)

		if !b.IsAlreadyShot(x, y) {
			return coordinate.NewCoordinate(x, y)
		}
	}
}
