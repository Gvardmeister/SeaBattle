package game

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/interfaces"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockboard"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockplayer"
)

func TestGameStart_Original_StartMethodCovered(t *testing.T) {
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			return true
		},
	}
	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			return false
		},
	}

	player1 := &mockplayer.MockPlayer{
		Name: "Player1",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(0, 0), false
		},
	}
	player2 := &mockplayer.MockPlayer{
		Name: "Player2",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(1, 1), false
		},
	}

	g := &game{
		player1: player1,
		board1:  board1,
		player2: player2,
		board2:  board2,
	}

	g.Start()

	if !board1.PlaceShipsCalled {
		t.Errorf("board1.PlaceShips() не вызван")
	}
	if !board2.PlaceShipsCalled {
		t.Errorf("board2.PlaceShips() не вызван")
	}
}

// func TestStart_Player1WinsWhenBoard2AllShipsKill(t *testing.T) {
// 	mBoard1 := &mockboard.MockBoard{}

// 	mBoard2 := &mockboard.MockBoard{
// 		AllShipsKillFunc: func() bool {
// 			return true
// 		},
// 	}

// 	mplayer1 := &mockplayer.MockPlayer{
// 		Name: "Player1",
// 		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
// 			return coordinate.NewCoordinate(0, 0), false
// 		},
// 	}
// 	mplayer2 := &mockplayer.MockPlayer{
// 		Name: "Player2",
// 		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
// 			return coordinate.NewCoordinate(1, 1), false
// 		},
// 	}

// 	turnCalls := 0
// 	tg := &testGameStr{
// 		game: game{
// 			board1:  mBoard1,
// 			board2:  mBoard2,
// 			player1: mplayer1,
// 			player2: mplayer2,
// 		},
// 		turnFunc: func(p interfaces.Player, enemyBoard interfaces.Board) bool {
// 			turnCalls++
// 			return false
// 		},
// 	}

// 	tg.Start()

// 	if turnCalls < 1 {
// 		t.Errorf("Ожидался хотя бы 1 ход от player1, но было %d", turnCalls)
// 	}
// }
