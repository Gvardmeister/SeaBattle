package game

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/interfaces"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockboard"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockplayer"
)

func TestStart_CallsPlaceShips(t *testing.T) {
	mBoard1 := &mockboard.MockBoard{}
	mBoard2 := &mockboard.MockBoard{}

	mplayer1 := &mockplayer.MockPlayer{Name: "Player1"}
	mplayer2 := &mockplayer.MockPlayer{Name: "Player2"}

	game := &game{
		board1:  mBoard1,
		board2:  mBoard2,
		player1: mplayer1,
		player2: mplayer2,
	}

	tg := &testGameStr{
		game: *game,
		turnFunc: func(p interfaces.Player, enemyBoard interfaces.Board) bool {
			return true
		},
	}

	tg.Start()

	if !mBoard1.PlaceShipsCalled {
		t.Errorf("PlaceShips не был вызван для board1")
	}
	if !mBoard2.PlaceShipsCalled {
		t.Errorf("PlaceShips не был вызван для board2")
	}
}
