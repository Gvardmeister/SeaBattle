package main

import (
	"fmt"
)

const (
	deck   = "X"
	loss   = "."
	change = "@"
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
	cells coordinate
	hit   []bool
}

type coordinate struct {
	x int
	y int
}

func initGame() {

}

func main() {
	fmt.Println()

	initGame()
}
