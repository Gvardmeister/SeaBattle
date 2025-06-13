package board

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/domain/ship"
)

const (
	Deck       = "X"
	Loss       = "."
	Change     = "@"
	SizeBoard  = 10
	Horizontal = 0
	Vertical   = 1
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

		if b.сanPlaceShip(shipX, shipY, positionShip, sizeShip) {
			b.placeShip(shipX, shipY, positionShip, sizeShip)
			break
		}
	}
}

func (b *Board) сanPlaceShip(x, y, orientation, sizeShip int) bool {
	dx, dy := 0, 0

	if orientation == Horizontal {
		dy = 1
	} else if orientation == Vertical {
		dx = 1
	}

	for i := 0; i < sizeShip; i++ {
		nx := x + dx*i
		ny := y + dy*i

		if IsOutOfBounds(nx, ny) {
			return false
		}

		if b.hasNeighboringShips(nx, ny) {
			return false
		}
	}

	return true
}

func IsOutOfBounds(x, y int) bool { return x < 0 || y < 0 || x >= SizeBoard || y >= SizeBoard }

func (b *Board) hasNeighboringShips(x, y int) bool {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			cx := x + dx
			cy := y + dy

			if !IsOutOfBounds(cx, cy) {
				continue
			}

			for _, ship := range b.Ships {
				for _, c := range ship.Cells {
					if c.GetX() == cx && c.GetY() == cy {
						return true
					}
				}
			}
		}
	}

	return false
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

	if orientation == Horizontal {
		dy = 1
	} else if orientation == Vertical {
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

func (b *Board) IsAlreadyShot(x, y int) bool {
	return b.Grid[x][y] != Change
}
