package board

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/domain/ship"
)

const (
	Deck      = "X"
	Loss      = "."
	Change    = "@"
	SizeBoard = 10
)

type Board struct {
	Grid  [][]string
	Ships []ship.Ship
}

func NewBoard(sizeBoard int) *Board {
	grid := make([][]string, sizeBoard)

	for i := 0; i < sizeBoard; i++ {
		grid[i] = make([]string, sizeBoard)
		for j := 0; j < sizeBoard; j++ {
			grid[i][j] = Change
		}
	}

	return &Board{
		Grid:  grid,
		Ships: []ship.Ship{},
	}
}

func (b *Board) Render() {
	fmt.Println("   1 2 3 4 5 6 7 8 9 10")

	for i := 0; i < SizeBoard; i++ {
		fmt.Printf("%2d ", i+1)
		for j := 0; j < SizeBoard; j++ {
			fmt.Print(b.Grid[i][j], " ")
		}
		fmt.Println()
	}
}

func (b *Board) generateShip(sizeShip int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		shipX := generator.Intn(SizeBoard)
		shipY := generator.Intn(SizeBoard)
		positionShip := generator.Intn(2)

		if b.canPlaceShip(shipX, shipY, positionShip, sizeShip) {
			b.placeShip(shipX, shipY, positionShip, sizeShip)
			break
		}
	}
}

func (b *Board) canPlaceShip(x, y, orientation, sizeShip int) bool {
	dx, dy := 0, 0
	if orientation == 0 {
		dy = 1
	} else {
		dx = 1
	}

	for i := 0; i < sizeShip; i++ {
		nx := x + dx*i
		ny := y + dy*i

		if nx < 0 || ny < 0 || nx >= SizeBoard || ny >= SizeBoard {
			return false
		}

		for dx2 := -1; dx2 <= 1; dx2++ {
			for dy2 := -1; dy2 <= 1; dy2++ {
				cx := nx + dx2
				cy := ny + dy2

				if cx >= 0 && cy >= 0 && cx < SizeBoard && cy < SizeBoard {
					for _, ship := range b.Ships {
						for _, c := range ship.Cells {
							if c.GetX() == cx && c.GetY() == cy {
								return false
							}
						}
					}
				}
			}
		}
	}

	return true
}

func (b *Board) PlaceShips() {
	variationShip := [][]int{
		{4, 1},
		{3, 2},
		{2, 3},
		{1, 4},
	}

	for _, value := range variationShip {
		sizeShip := value[0]
		count := value[1]

		for i := 0; i < count; i++ {
			b.generateShip(sizeShip)
		}
	}
}

func (b *Board) placeShip(x, y, orientation, sizeShip int) {
	dx, dy := 0, 0
	if orientation == 0 {
		dy = 1
	} else {
		dx = 1
	}

	var cells []coordinate.Coordinate

	for i := 0; i < sizeShip; i++ {
		nx := x + dx*i
		ny := y + dy*i
		cells = append(cells, coordinate.NewCoordinate(nx, ny))
	}

	b.Ships = append(b.Ships, *ship.NewShip(cells))
}

func (b *Board) AllShipsKill() bool {
	for _, ship := range b.Ships {
		for _, partHit := range ship.Hit {
			if !partHit {
				return false
			}
		}
	}
	return true
}
