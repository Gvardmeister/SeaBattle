package main

import (
	"fmt"
)

// const coordinate_Y = 1 // Вроде нельзя передать констату в структуру

type ship struct {
	sizeShip     int
	coordinate_X []int
	coordinate_Y int // поле всегда будет 1, т.к. пока что это константа
	countShip    int
}

type mapBattale struct {
	sizeX []int // 10
	sizeY int   // 1
}

func (mb *mapBattale) newMapBattle(newSizeX []int) string {
	(*mb).sizeX = newSizeX

	return ""
}

func gameBattle() {
	shipObj := ship{
		sizeShip:     4,
		coordinate_X: []int{5, 6, 7, 8},
		coordinate_Y: 1,
		countShip:    1,
	}

	mapBattaleObj := mapBattale{
		sizeX: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		sizeY: 1,
	}

	var pointerMapBattle *mapBattale = &mapBattaleObj
	pointerMapBattle.newMapBattle([]int{}) // проверял теорию заглушка

	fmt.Println(shipObj, pointerMapBattle) // заглушка от ошибки
}

func main() {
	fmt.Println("@@@@@@@@@@")

	var target int

	fmt.Print("Введите координату: ")
	fmt.Scan(&target)

	gameBattle()
}
