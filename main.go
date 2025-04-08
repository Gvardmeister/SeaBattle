package main

import (
	"fmt"
	"strconv"
)

func main() {
	var target string

	count := 0
	ship := []string{"@", "@", "@", "@", "@", "@", "@", "@", "@", "@"}
	mapBattle := map[string]int{
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
	}

	for {
		fmt.Println("\nВведите координату: ")
		fmt.Scan(&target)

		if val, ok := mapBattle[target]; ok {
			fmt.Println("Вы попали")

			count += 1
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
			fmt.Println("Вы выиграли!")
			break
		}
	}
}
