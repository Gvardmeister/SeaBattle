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

func generationShip(size int, ship [][]bool, sizeShip int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	shipX := generator.Intn(size - sizeShip)
	shipY := generator.Intn(size - sizeShip)
	positionShip := generator.Intn(2)

	switchPositionShip(shipX, shipY, positionShip, sizeShip, ship)
}

func switchPositionShip(shipX, shipY, positionShip, sizeShip int, ship [][]bool) {
	switch positionShip {
	case 0:
		for i := 0; i < sizeShip; i++ {
			ship[shipX][shipY+i] = true
		}
	case 1:
		for i := 0; i < sizeShip; i++ {
			ship[shipX+i][shipY] = true
		}
	}
}

func printShip(size int, ship [][]bool) {
	generationShip(size, ship, 4)
	generationShip(size, ship, 3)
	// generationShip(size, ship, 3)
	generationShip(size, ship, 2)
	// generationShip(size, ship, 2)
	// generationShip(size, ship, 2)
	generationShip(size, ship, 1)
	// generationShip(size, ship, 1)
	// generationShip(size, ship, 1)
	// generationShip(size, ship, 1)
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

	printShip(size, ship)

	// убрать потом 2 цикла
	for deadCount < 20 {
		for i := 0; i <= 10; i++ {
			for j := 0; j <= 10; j++ {
				coordinateX = i
				coordinateY = j

				// ниже все вытащить из цикла
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
	}
	// fmt.Println("\nВведите координаты в формате {X и Y} от 1 до 10:")
	// fmt.Scan(&coordinateX, &coordinateY)
}
