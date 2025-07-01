package game

import (
	"github.com/Gvardmeister/seabattle/internal/interfaces"
)

type testGameStr struct {
	game
	turnFunc func(p interfaces.Player, enemyBoard interfaces.Board) bool
}

func (tg *testGameStr) Start() {
	tg.board1.PlaceShips()
	tg.board2.PlaceShips()

	for {
		if tg.turnFunc != nil {
			if tg.turnFunc(tg.player1, tg.board2) {
				break
			}
			if tg.turnFunc(tg.player2, tg.board1) {
				break
			}
		} else {
			if tg.turn(tg.player1, tg.board2) {
				break
			}
			if tg.turn(tg.player2, tg.board1) {
				break
			}
		}
	}
}
