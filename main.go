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
	cells [][2]int // можно было разбить на еще одну структуру, чтобы выделить только координаты
	hit   []bool
}

func initGame() {

}

func main() {
	fmt.Println()

	initGame()
}
