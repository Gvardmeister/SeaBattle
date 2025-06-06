package main

import (
	"github.com/Gvardmeister/seabattle/internal/domain/game"
	"github.com/Gvardmeister/seabattle/internal/services"
)

func main() {
	inputReader := services.NewStdInputReader()
	game.NewGame(inputReader).Start()
}
