package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	ship   = "X"
	loss   = "."
	change = "@"
)

func plusNumber(number int, MB map[string]int) {
	countMap := 1
	strKey := strconv.Itoa(number)
	MB[strKey] = number

	for countMap != 4 {
		number++
		strKey := strconv.Itoa(number)
		MB[strKey] = number
		countMap++
	}
}

func minusNumber(number int, MB map[string]int) {
	countMap := 1
	strKey := strconv.Itoa(number)
	MB[strKey] = number

	for countMap != 4 {
		number--
		strKey := strconv.Itoa(number)
		MB[strKey] = number
		countMap++
	}
}

func main() {
	deadCount := 0

	var target string
	mapBattle := make(map[string]int)
	boar := make([]string, 10)

	for i := range boar {
		boar[i] = change
	}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberGenerator := generator.Intn(11)

	if numberGenerator >= 1 && numberGenerator <= 7 {
		plusNumber(numberGenerator, mapBattle)
	} else {
		minusNumber(numberGenerator, mapBattle)
	}

	for {
		fmt.Println("\nВведите координату: ")
		fmt.Scan(&target)

		_, err := strconv.Atoi(target)

		if err != nil {
			fmt.Println("Введите числовое значение!")
			continue
		}

		if _, ok := mapBattle[target]; ok {
			deadCount++

			if deadCount == 4 {
				fmt.Println("\nВы выиграли!")
			} else {
				fmt.Println("Вы попали")
			}

			for idx := range boar {
				if target == strconv.Itoa(idx+1) {
					boar[idx] = ship
				}
				fmt.Print(boar[idx])
			}
			fmt.Println()
		} else {
			fmt.Println("Вы промахнулись")

			for idx := range boar {
				if target == strconv.Itoa(idx+1) {
					boar[idx] = loss
				}
				fmt.Print(boar[idx])
			}
			fmt.Println()
		}

		if deadCount == 4 {
			break
		}
	}
}
