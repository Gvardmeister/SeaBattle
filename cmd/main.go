package main

import (
	"github.com/Gvardmeister/seabattle/internal/domain/game"
	"github.com/Gvardmeister/seabattle/internal/services/stdinputreader"
)

func main() {
	inputReader := stdinputreader.NewStdInputReader()
	game.NewGame(inputReader).Start()
}
