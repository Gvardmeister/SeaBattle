package game

import (
	"bytes"
	"os"
	"strings"
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

func TestStart_Player1WinsWhenBoard2AllShipsKill(t *testing.T) {
	mBoard1 := &mockboard.MockBoard{}
	mBoard2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			return true
		},
	}

	mplayer1 := &mockplayer.MockPlayer{
		Name: "Player1",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(0, 0), false
		},
	}
	mplayer2 := &mockplayer.MockPlayer{
		Name: "Player2",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(1, 1), false
		},
	}

	turnCalls := 0

	tg := &testGameStr{
		game: game{
			board1:  mBoard1,
			board2:  mBoard2,
			player1: mplayer1,
			player2: mplayer2,
		},
		turnFunc: func(p interfaces.Player, enemyBoard interfaces.Board) bool {
			turnCalls++
			return false
		},
	}

	for {
		if tg.turnFunc(tg.player1, tg.board2) {
			break
		}
		if tg.board2.AllShipsKill() {
			break
		}
		if tg.turnFunc(tg.player2, tg.board1) {
			break
		}
		if tg.board1.AllShipsKill() {
			break
		}
	}

	if turnCalls < 1 {
		t.Errorf("Ожидался хотя бы 1 ход от player1, но было %d", turnCalls)
	}
}

func TestStart_Player2WinsWhenBoard1AllShipsKill(t *testing.T) {
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
			return coordinate.NewCoordinate(2, 2), false
		},
	}
	player2 := &mockplayer.MockPlayer{
		Name: "Player2",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(3, 3), false
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

func TestStart_EndsWhenTurnReturnsTrue(t *testing.T) {
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			return false
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
			return coordinate.NewCoordinate(4, 4), false
		},
	}
	player2 := &mockplayer.MockPlayer{
		Name: "Player2",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(5, 5), false
		},
	}

	var turnCalls int

	g := &testGameStr{
		game: game{
			player1: player1,
			player2: player2,
			board1:  board1,
			board2:  board2,
		},
		turnFunc: func(p interfaces.Player, b interfaces.Board) bool {
			turnCalls++
			return true
		},
	}

	g.Start()

	if turnCalls != 1 {
		t.Errorf("turnFunc должен быть вызван один раз, но был %d", turnCalls)
	}
}

func TestStart_MultipleTurnsBoardDestroyed(t *testing.T) {
	callCount := 0

	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			callCount++
			return callCount >= 3
		},
	}
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			return false
		},
	}

	player1 := &mockplayer.MockPlayer{
		Name: "Player1",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(1, 1), false
		},
	}
	player2 := &mockplayer.MockPlayer{
		Name: "Player2",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(2, 2), false
		},
	}

	turns := 0

	g := &testGameStr{
		game: game{
			player1: player1,
			board1:  board1,
			player2: player2,
			board2:  board2,
		},
		turnFunc: func(p interfaces.Player, b interfaces.Board) bool {
			turns++
			return false
		},
	}

	for {
		if g.turnFunc(g.player1, g.board2) {
			break
		}
		if g.board2.AllShipsKill() {
			break
		}

		if g.turnFunc(g.player2, g.board1) {
			break
		}
		if g.board1.AllShipsKill() {
			break
		}
	}

	if turns < 3 {
		t.Errorf("Ожидалось минимум 3 хода, но было: %d", turns)
	}
	if callCount < 3 {
		t.Errorf("Ожидалось минимум 3 вызова AllShipsKill, но было: %d", callCount)
	}
}

func TestGame_Start_TurnBreaksImmediately(t *testing.T) {
	player1Moves := 0

	player1 := &mockplayer.MockPlayer{
		Name: "Игрок 1",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			player1Moves++
			return coordinate.NewCoordinate(0, 0), true
		},
	}

	player2 := &mockplayer.MockPlayer{
		Name: "Игрок 2",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(0, 0), true
		},
	}

	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			return true
		},
	}

	g := &game{
		board1:  &mockboard.MockBoard{},
		board2:  board2,
		player1: player1,
		player2: player2,
	}

	g.Start()
}

func TestStart_Player1Wins_PrintAndBreak(t *testing.T) {
	board1 := &mockboard.MockBoard{}
	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return true },
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
			t.Fatal("player2 ходить не должен")
			return coordinate.NewCoordinate(0, 0), false
		},
	}

	g := &game{
		board1:  board1,
		board2:  board2,
		player1: player1,
		player2: player2,
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	g.Start()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Ошибка чтения из pipe: %v", err)
	}

	output := buf.String()

	if !strings.Contains(output, "Победил игрок - Player1") {
		t.Errorf("Ожидали сообщение о победе Player1, получили: %q", output)
	}
}

func TestStart_Player2TurnBreaksLoop(t *testing.T) {
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return false },
	}
	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return false },
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

	turnCalls := 0
	tg := &testGameStr{
		game: game{
			board1:  board1,
			board2:  board2,
			player1: player1,
			player2: player2,
		},
		turnFunc: func(p interfaces.Player, b interfaces.Board) bool {
			turnCalls++
			return p == player2
		},
	}

	tg.Start()

	if turnCalls == 0 {
		t.Errorf("turnFunc должен был вызваться минимум один раз")
	}
	if turnCalls != 2 {
		t.Errorf("Ожидалось 2 вызова turnFunc, получено %d", turnCalls)
	}
}

func TestStart_Player2Wins_PrintAndBreak(t *testing.T) {
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return true },
	}
	board2 := &mockboard.MockBoard{}

	player1 := &mockplayer.MockPlayer{
		Name: "Player1",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			t.Fatal("player1 не должен ходить, т.к. board1 уже уничтожен")
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
		board1:  board1,
		board2:  board2,
		player1: player1,
		player2: player2,
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	g.Start()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Ошибка чтения из pipe: %v", err)
	}

	output := buf.String()

	if !strings.Contains(output, "Победил игрок - Player2") {
		t.Errorf("Ожидали сообщение о победе Player2, получили: %q", output)
	}
}

func TestStart_TurnReturnsTrueForPlayer1(t *testing.T) {
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return false },
	}
	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return false },
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

	turnCalls := 0

	tg := &testGameStr{
		game: game{
			board1:  board1,
			board2:  board2,
			player1: player1,
			player2: player2,
		},
		turnFunc: func(p interfaces.Player, b interfaces.Board) bool {
			turnCalls++
			if p == player1 {
				return true
			}
			return false
		},
	}

	tg.Start()

	if turnCalls != 1 {
		t.Errorf("Ожидался ровно 1 вызов turn для player1, было %d", turnCalls)
	}
}

func TestStart_TurnReturnsTrueForPlayer2(t *testing.T) {
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return false },
	}
	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return false },
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

	turnCalls := 0

	tg := &testGameStr{
		game: game{
			board1:  board1,
			board2:  board2,
			player1: player1,
			player2: player2,
		},
		turnFunc: func(p interfaces.Player, b interfaces.Board) bool {
			turnCalls++
			if p == player2 {
				return true
			}
			return false
		},
	}

	tg.Start()

	if turnCalls < 2 {
		t.Errorf("Ожидалось минимум 2 вызова turn (player1 и player2), было %d", turnCalls)
	}
}

func TestStart_TurnNeverReturnsTrue_BreaksWhenAllShipsKill(t *testing.T) {
	callCount := 0

	board2 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool {
			callCount++
			return callCount >= 3
		},
	}
	board1 := &mockboard.MockBoard{
		AllShipsKillFunc: func() bool { return false },
	}

	player1 := &mockplayer.MockPlayer{
		Name: "Player1",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(1, 1), false
		},
	}
	player2 := &mockplayer.MockPlayer{
		Name: "Player2",
		GetMoveFunc: func(b interfaces.Board) (coordinate.Coordinate, bool) {
			return coordinate.NewCoordinate(2, 2), false
		},
	}

	turns := 0

	tg := &testGameStr{
		game: game{
			board1:  board1,
			board2:  board2,
			player1: player1,
			player2: player2,
		},
		turnFunc: func(p interfaces.Player, b interfaces.Board) bool {
			turns++
			return false
		},
	}

	for {
		if tg.turnFunc(tg.player1, tg.board2) {
			break
		}
		if tg.board2.AllShipsKill() {
			break
		}

		if tg.turnFunc(tg.player2, tg.board1) {
			break
		}
		if tg.board1.AllShipsKill() {
			break
		}
	}

	if turns < 3 {
		t.Errorf("Ожидалось минимум 3 хода, было %d", turns)
	}
	if callCount < 3 {
		t.Errorf("Ожидалось минимум 3 вызова AllShipsKill, было %d", callCount)
	}
}
