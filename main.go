package main

import (
	"fmt"
)

const (
	deck      = "X"
	loss      = "."
	change    = "@"
	sizeBoard = 10
)

type game struct {
	player player
	board  board
}

type player struct {
	position coordinate
}

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
	board := newBoard(sizeBoard)
	player := newPlayer()

	return &game{
		player: *player,
		board:  *board,
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

func newPlayer() *player {
	var coord coordinate
	coord.PromptCoordinate()

	return &player{
		position: coord,
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

func (c *coordinate) Get() (int, int) {
	return c.xcoordinate, c.ycoordinate
}

func (c *coordinate) PromptCoordinate() {
	for {
		var x, y int
		fmt.Println("Введите координаты X и Y {от 1 до 10}:")
		fmt.Scan(&x, &y)

		if x >= 1 && x <= sizeBoard && y >= 1 && y <= sizeBoard {
			c.xcoordinate = x - 1
			c.ycoordinate = y - 1
			break
		}
		fmt.Println("Неверные координаты. Повторите ввод.")
	}
}

func initGame() {
}

func main() {
	initGame()
}
