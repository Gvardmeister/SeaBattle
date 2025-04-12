package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var target string
	mapBattle := make(map[string]int)

	count := 0
	countMap := 0
	ship := []string{"@", "@", "@", "@", "@", "@", "@", "@", "@", "@"}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		numberGenerator := generator.Intn(10) + 1
		strKey := strconv.Itoa(numberGenerator)

		if numberGenerator, ok := mapBattle[strKey]; ok {
			continue
		} else {
			mapBattle[strKey] = numberGenerator
			countMap++
		}

		if countMap == 4 {
			break
		}
		// чтобы числа шли подряд
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
			count++

			if count == 4 {
				fmt.Println("\nВы выиграли!")
			} else {
				fmt.Println("Вы попали")
			}

			for idx := range ship {
				if target == strconv.Itoa(idx+1) {
					ship[idx] = "X"
				}
				fmt.Print(ship[idx])
			}
			fmt.Println()
		} else {
			fmt.Println("Вы промахнулись")

			for idx := range ship {
				if target == strconv.Itoa(idx+1) {
					ship[idx] = "."
				}
				fmt.Print(ship[idx])
			}
			fmt.Println()
		}

		if count == 4 {
			break
		}
	}
}
