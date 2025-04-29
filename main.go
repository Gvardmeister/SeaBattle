package main

import (
	"fmt"
	"math/rand"
	"time"
)

// реализовать квадратное поле
// реализовать генерацию короблей

const (
	deck   = "X"
	loss   = "."
	change = "@"
)

func main() {
	deadCount := 0

	var target int
	ship := make([]bool, 10)
	boar := make([][]string, 10)

	for i := 0; i < 10; i++ {
		boar[i] = make([]string, 10)

		for j := 0; j < 10; j++ {
			boar[i][j] = change
		}
	}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberGenerator := generator.Intn(7)

	for i := 0; i < 4; i++ {
		ship[numberGenerator+i] = true
	}

	for {
		fmt.Println("\nВведите координату: ")
		fmt.Scan(&target)
		fmt.Println()

		if ship[target-1] == true {
			deadCount++

			if deadCount == 4 {
				fmt.Println("Вы выиграли")
			} else {
				fmt.Println("Вы попали")
			}

			for i := range boar {
				if target == i+1 {
					boar[i] = deck
				}
				fmt.Print(boar[i])
			}
			fmt.Println()
		} else {
			fmt.Println("Вы промахнулись")

			for i := range boar {
				if target == i+1 {
					boar[i] = loss
				}
				fmt.Print(boar[i])
			}
			fmt.Println()
		}

		if deadCount == 4 {
			break
		}
	}
}
