package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	deck      = "X"
	loss      = "."
	change    = "@"
	sizeBoard = 10
)

type Player interface {
	GetMove(grid [][]string) coordinate
}

type game struct {
	player Player
	board  board
}

type HumanPlayer struct{}

type board struct {
	grid  [][]string
	ships []ship
}

type ship struct {
	cells []coordinate
	hit   []bool
}

type coordinate struct {
	xcoordinate int
	ycoordinate int
}

func newGame() *game {
	return &game{
		player: &HumanPlayer{},
		board:  *newBoard(sizeBoard),
	}
}

func newBoard(sizeBoard int) *board {
	grid := make([][]string, sizeBoard)

	for i := 0; i < sizeBoard; i++ {
		grid[i] = make([]string, sizeBoard)
		for j := 0; j < sizeBoard; j++ {
			grid[i][j] = change
		}
	}

	return &board{
		grid:  grid,
		ships: []ship{},
	}
}

func newCoordinate(x, y int) coordinate {
	return coordinate{
		xcoordinate: x,
		ycoordinate: y,
	}
}

func newShip(cells []coordinate) *ship {
	return &ship{
		cells: cells,
		hit:   make([]bool, len(cells)),
	}
}

func (hp *HumanPlayer) GetMove(grid [][]string) coordinate {
	for {
		var x, y int
		fmt.Println("\nВведите координаты X и Y {от 1 до 10}:")
		fmt.Scan(&x, &y)

		if x < 1 || x > sizeBoard || y < 1 || y > sizeBoard {
			fmt.Println("\nНеверные координаты. Повторите ввод.")
			continue
		}

		x--
		y--

		if grid[y][x] == change {
			return newCoordinate(x, y)
		}
		fmt.Println("\nВы уже стреляли в эту клетку.")
	}
}

func (c *coordinate) Get() (int, int) {
	return c.xcoordinate, c.ycoordinate
}

func (b *board) Render() {
	fmt.Println("   1 2 3 4 5 6 7 8 9 10")

	for i := 0; i < sizeBoard; i++ {
		fmt.Printf("%2d ", i+1)
		for j := 0; j < sizeBoard; j++ {
			fmt.Print(b.grid[i][j], " ")
		}
		fmt.Println()
	}
}

func (b *board) generateShip(sizeShip int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		shipX := generator.Intn(sizeBoard)
		shipY := generator.Intn(sizeBoard)
		positionShip := generator.Intn(2)

		if b.canPlaceShip(shipX, shipY, positionShip, sizeShip) {
			b.placeShip(shipX, shipY, positionShip, sizeShip)
			break
		}
	}
}

func (b *board) canPlaceShip(x, y, orientation, sizeShip int) bool {
	dx, dy := 0, 0
	if orientation == 0 {
		dy = 1
	} else {
		dx = 1
	}

	for i := 0; i < sizeShip; i++ {
		nx := x + dx*i
		ny := y + dy*i

		if nx < 0 || ny < 0 || nx >= sizeBoard || ny >= sizeBoard {
			return false
		}

		for dx2 := -1; dx2 <= 1; dx2++ {
			for dy2 := -1; dy2 <= 1; dy2++ {
				cx := nx + dx2
				cy := ny + dy2

				if cx >= 0 && cy >= 0 && cx < sizeBoard && cy < sizeBoard {
					for _, ship := range b.ships {
						for _, c := range ship.cells {
							if c.xcoordinate == cx && c.ycoordinate == cy {
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

func (b *board) PlaceShips() {
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

func (b *board) placeShip(x, y, orientation, sizeShip int) {
	dx, dy := 0, 0
	if orientation == 0 {
		dy = 1
	} else {
		dx = 1
	}

	var cells []coordinate

	for i := 0; i < sizeShip; i++ {
		nx := x + dx*i
		ny := y + dy*i
		cells = append(cells, newCoordinate(nx, ny))
	}

	b.ships = append(b.ships, *newShip(cells))
}

func (g *game) MakeMove(coord coordinate) {
	x, y := coord.xcoordinate, coord.ycoordinate
	hit := false

	for i := range g.board.ships {
		ship := &g.board.ships[i]

		for j, cell := range ship.cells {
			if cell.xcoordinate == x && cell.ycoordinate == y {
				ship.hit[j] = true
				hit = true

				break
			}
		}

		if hit {
			break
		}
	}

	if hit {
		g.board.grid[x][y] = deck

		fmt.Println("\nПопадание!")
	} else {
		g.board.grid[x][y] = loss

		fmt.Println("\nВы промахнулись!")
	}

	fmt.Println()
	g.board.Render()
}

func (g *game) initGame() {
	g.board.PlaceShips()
	g.board.Render()

	for {
		coord := g.player.GetMove(g.board.grid)
		g.MakeMove(coord)
	}
}

func main() {
	newGame().initGame()
}
