package bot

import (
	"testing"

	"github.com/Gvardmeister/seabattle/internal/domain/coordinate"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockboard"
	"github.com/Gvardmeister/seabattle/internal/test/mocks/mockcoordinategenerator"
)

func TestNewBotPlayer(t *testing.T) {
	mockGen := &mockcoordinategenerator.MockCoordinateGenerator{
		FixedCoord: coordinate.NewCoordinate(2, 9),
	}

	mockBoard := &mockboard.MockBoard{
		IsAlreadyShotFunc: func(x, y int) bool { return false },
	}

	bot := NewBotPlayer("Bot", mockGen)

	coord, _ := bot.GetMove(mockBoard)

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
