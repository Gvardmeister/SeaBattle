package game_test

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/game"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockinputreader"
)

func TestNewGame_PvPMode(t *testing.T) {
	input := mockinputreader.NewMockInput([]string{"1", "Alex", "Nasty"})
	g := game.NewGame(input)

	if g.GetPlayer1().GetName() != "Alex" {
		t.Errorf("Ожидалось имя игрока 'Alex', получено %s", g.GetPlayer1().GetName())
	}

	if g.GetPlayer2().GetName() != "Nasty" {
		t.Errorf("Ожидалось имя игрока 'Nasty', получено %s", g.GetPlayer2().GetName())
	}
}

func TestNewGame_PvEMode(t *testing.T) {
	input := mockinputreader.NewMockInput([]string{"2", "Alex"})
	g := game.NewGame(input)

	if g.GetPlayer1().GetName() != "Alex" {
		t.Errorf("Ожидалось имя игрока 'Alex', получено %s", g.GetPlayer1().GetName())
	}

	if g.GetPlayer2().GetName() != "Bot" {
		t.Errorf("Ожидалось имя игрока 'Bot', получено %s", g.GetPlayer2().GetName())
	}
}

func TestNewGame_ReadLineError(t *testing.T) {
	input := mockinputreader.NewMockErrorInputReader()
	g := game.NewGame(input)

	if g.GetPlayer1().GetName() != "Player1" {
		t.Errorf("Ожидалось имя 'Player1', получено: %q", g.GetPlayer1().GetName())
	}

}

func TestNewGame_InvalidMode(t *testing.T) {
	input := mockinputreader.NewMockInput([]string{"asda", "Alex"})
	g := game.NewGame(input)

	if g.GetPlayer1().GetName() != "Alex" {
		t.Errorf("Ожидалось имя 'Alex', получено %s", g.GetPlayer1().GetName())
	}

	if g.GetPlayer2().GetName() != "Bot" {
		t.Errorf("Ожидался Бот', получен: %s", g.GetPlayer2().GetName())
	}
}

func TestNewGame_EmptyModeInput(t *testing.T) {
	input := mockinputreader.NewMockInput([]string{"", "Alex"})
	g := game.NewGame(input)

	if g.GetPlayer1().GetName() != "Alex" {
		t.Errorf("Ожидалось имя 'Alex', получено: %s", g.GetPlayer1().GetName())
	}

	if g.GetPlayer2().GetName() != "Bot" {
		t.Errorf("Ожидалось имя 'Bot', получено: %q", g.GetPlayer2().GetName())
	}
}

func TestNewGame_PvPMode_EmptyNames(t *testing.T) {
	input := mockinputreader.NewMockInput([]string{"1", "", ""})
	g := game.NewGame(input)

	if g.GetPlayer1().GetName() != "Player1" {
		t.Errorf("Ожидалось имя 'Player1', получено: %q", g.GetPlayer1().GetName())
	}

	if g.GetPlayer2().GetName() != "Player2" {
		t.Errorf("Ожидалось имя 'Player2', получено: %q", g.GetPlayer2().GetName())
	}
}

func TestNewGame_PvEMode_EmptyName(t *testing.T) {
	input := mockinputreader.NewMockInput([]string{"2", ""})
	g := game.NewGame(input)

	if g.GetPlayer1().GetName() != "Player1" {
		t.Errorf("Ожидалось имя 'Player1', получено: %q", g.GetPlayer1().GetName())
	}

	if g.GetPlayer2().GetName() != "Bot" {
		t.Errorf("Ожидалось имя бота 'Bot', получено: %q", g.GetPlayer2().GetName())
	}
}
