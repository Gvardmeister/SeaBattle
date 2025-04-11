package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var mapBattle map[string]int // тестовое значение мапы
	var target string

	count := 0
	ship := []string{"@", "@", "@", "@", "@", "@", "@", "@", "@", "@"}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	countGenerator := 4
	for i := 0; i < countGenerator; i++ {
		numberGenerator := generator.Intn(10) + 1

		strKey := strconv.Itoa(numberGenerator)
		mapBattle[strKey] = numberGenerator
	}

	// mapBattle := map[string]int{
	// 	"5": 5,
	// 	"6": 6,
	// 	"7": 7,
	// 	"8": 8,
	// }

	for {
		fmt.Println("\nВведите координату: ")
		fmt.Scan(&target)

		_, err := strconv.Atoi(target)

		if err != nil {
			fmt.Println("Введите числовое значение!")
			continue
		}

		if val, ok := mapBattle[target]; ok {
			count += 1

			if count == 4 {
				fmt.Println("\nВы выиграли!")
			} else {
				fmt.Println("Вы попали")
			}

			for idx := range ship {
				if val == idx+1 {
					ship[idx] = "X"
				}
				fmt.Print(ship[idx])
			}
		} else {
			fmt.Println("Вы промахнулись")

			for idx := range ship {
				if target == strconv.Itoa(idx+1) {
					ship[idx] = "."
				}
				fmt.Print(ship[idx])
			}
		}

		if count == 4 {
			break
		}
	}
}

// Генерация рандомных чисел в мапу + проверка
