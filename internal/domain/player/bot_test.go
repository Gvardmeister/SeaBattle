package player

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/board"
	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/test/mocks"
)

func TestNewBotPlayer(t *testing.T) {
	mockGen := &mocks.MockCoordinateGenerator{
		FixedCoord: coordinate.NewCoordinate(2, 9),
	}

	bot := NewBotPlayer("Bot", mockGen)
	board := &board.Board{}

	coord, _ := bot.GetMove(board)

	if coord.GetX() != 2 || coord.GetY() != 9 {
		t.Errorf("Ожидались координаты (2, 9), получено (%d, %d)", coord.GetX(), coord.GetY())
	}

	if !mockGen.Called {
		t.Errorf("Ожидалось, что будет вызван Generator")
	}

	if bot.GetName() != "Bot" {
		t.Errorf("Ожидалось имя 'Bot', получено %s", bot.GetName())
	}
}
