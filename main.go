package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	deck   = "X"
	loss   = "."
	change = "@"
)

func printBoar(size int, boar [][]string) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(boar[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	var coordinateX, coordinateY int
	deadCount := 0
	size := 10

	boar := make([][]string, size)
	for i := 0; i < size; i++ {
		boar[i] = make([]string, size)

		for j := 0; j < size; j++ {
			boar[i][j] = change
		}
	}

	ship := make([][]bool, size)
	for i := range ship {
		ship[i] = make([]bool, size)
	}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	shipX := generator.Intn(size - 4)
	shipY := generator.Intn(size - 4)
	positionShip := generator.Intn(2)

	switch positionShip {
	case 0:
		for i := 0; i < 4; i++ {
			ship[shipX][shipY+i] = true
		}
	case 1:
		for i := 0; i < 4; i++ {
			ship[shipX+i][shipY] = true
		}
	}

	for deadCount < 4 {
		fmt.Println("\nВведите координаты в формате {X и Y} от 1 до 10:")
		fmt.Scan(&coordinateX, &coordinateY)
		fmt.Println()

		if coordinateX < 1 || coordinateX > size || coordinateY < 1 || coordinateY > size {
			fmt.Println("Неверные координаты")

			continue
		}

		if ship[coordinateX-1][coordinateY-1] == true && boar[coordinateX-1][coordinateY-1] != deck {
			boar[coordinateX-1][coordinateY-1] = deck
			deadCount++

			switch deadCount {
			case 4:
				fmt.Println("Вы выиграли")
			default:
				fmt.Println("Вы попали")
			}
		} else {
			fmt.Println("Вы промахнулись")

			boar[coordinateX-1][coordinateY-1] = loss
		}

		printBoar(size, boar)
	}
}
