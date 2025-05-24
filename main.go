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
	coordinateX int
	coordinateY int
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
	return &player{}
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

func initGame() {

}

func main() {
	fmt.Println()

	initGame()
}
