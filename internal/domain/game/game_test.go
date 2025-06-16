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

	if g.GetPlayer1().GetName() != "" {
		t.Errorf("Ожидалось пустое имя, получено: %q", g.GetPlayer1().GetName())
	}
}
