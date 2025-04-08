package main

import (
	"fmt"
	"strconv"
)

func main() {
	var target string

	fmt.Print("Введите координату: ")
	fmt.Scan(&target)

	ship := []string{"@", "@", "@", "@", "@", "@", "@", "@", "@", "@"}

	mapBattle := map[string]int{
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
	}

	if val, ok := mapBattle[target]; ok {
		fmt.Println("Вы попали")
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
}

// заклить, чтобы спрашивал постоянно
// сохранить результат
