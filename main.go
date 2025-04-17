package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

func main() {
	var target string
	mapBattle := make(map[string]int)

	count := 0
	countMap := 1
	ship := []string{"@", "@", "@", "@", "@", "@", "@", "@", "@", "@"}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberGenerator := generator.Intn(10) + 1

	// тестирую этот кусок
	if numberGenerator >= 1 && numberGenerator <= 7 {
		plusNumber(numberGenerator, mapBattle)

		// strKey := strconv.Itoa(numberGenerator)
		// mapBattle[strKey] = numberGenerator
		//
		// for countMap != 4 {
		// numberGenerator++
		// strKey := strconv.Itoa(numberGenerator)
		// mapBattle[strKey] = numberGenerator
		// countMap++
		// }
	} else {
		strKey := strconv.Itoa(numberGenerator)
		mapBattle[strKey] = numberGenerator

		for countMap != 4 {
			numberGenerator--
			strKey := strconv.Itoa(numberGenerator)
			mapBattle[strKey] = numberGenerator
			countMap++
		}
	}

	// рабочий кусок, но с багнутой генерацией
	// for {
	// numberGenerator := generator.Intn(10) + 1
	// strKey := strconv.Itoa(numberGenerator)
	//
	// if numberGenerator, ok := mapBattle[strKey]; ok {
	// continue
	// } else {
	// mapBattle[strKey] = numberGenerator
	// countMap++
	// }
	//
	// if countMap == 4 {
	// break
	// }
	// }

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
