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

	checkPlace := false

	for !checkPlace {
		shipX := generator.Intn(size - sizeShip)
		shipY := generator.Intn(size - sizeShip)
		positionShip := generator.Intn(2)

		if placeShip(shipX, shipY, positionShip, sizeShip, size, ship) {
			printPositionShip(shipX, shipY, positionShip, sizeShip, ship)

			checkPlace = true
		}
	}
}

func placeShip(shipX, shipY, positionShip, sizeShip, size int, ship [][]bool) bool {
	x := 0
	y := 0

	if positionShip == 0 {
		y = 1
	} else {
		x = 1
	}

	for i := 0; i < sizeShip; i++ {
		currentX := shipX + x*i
		currentY := shipY + y*i

		if currentX < 0 || currentY < 0 || currentX >= size || currentY >= size {
			return false
		}

		for x2 := -1; x2 <= 1; x2++ {
			for y2 := -1; y2 <= 1; y2++ {
				x3 := currentX + x2
				y3 := currentY + y2

				if x3 >= 0 && y3 >= 0 && x3 < size && y3 < size {
					if ship[x3][y3] {
						return false
					}
				}
			}
		}
	}

	return true
}

func printPositionShip(shipX, shipY, positionShip, sizeShip int, ship [][]bool) {
	if positionShip == 0 {
		for i := 0; i < sizeShip; i++ {
			ship[shipX][shipY+i] = true
		}
	} else {
		for i := 0; i < sizeShip; i++ {
			ship[shipX+i][shipY] = true
		}
	}
}

func printShip(size int, ship [][]bool) {
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
			generationShip(size, ship, sizeShip)
		}
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

	printShip(size, ship)

	for deadCount < 20 {
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

			if deadCount == 20 {
				fmt.Println("Вы выиграли")
			} else {
				fmt.Println("Вы попали")
			}
		} else if ship[coordinateX-1][coordinateY-1] == true && boar[coordinateX-1][coordinateY-1] == deck {
			fmt.Println("Вы уже сюда попадали")

			continue
		} else {
			fmt.Println("Вы промахнулись")

			boar[coordinateX-1][coordinateY-1] = loss
		}

		printBoar(size, boar)
	}
}
