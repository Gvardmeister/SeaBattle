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

func main() {
	deadCount := 0

	var target int
	ship := make([]bool, 10)

	boar := make([]string, 10)
	for i := range boar {
		boar[i] = change
	}

	for {
		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		numberGenerator := generator.Intn(11)

		fmt.Println("\nВведите координату: ")
		fmt.Scan(&target)
		fmt.Println()

		if target == numberGenerator {
			deadCount++

			if deadCount == 4 {
				fmt.Println("\nВы выиграли")
			} else {
				fmt.Println("Вы попали")
			}

			for i := range ship {
				ship[i] = true
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
