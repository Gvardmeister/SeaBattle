package coordinate

import "testing"

func TestNewCoordinate(t *testing.T) {
	cordinate := NewCoordinate(3, 7)

	if cordinate.GetX() != 3 {
		t.Errorf("Ожидалось х = 3, получено %d", cordinate.GetX())
	}

	if cordinate.GetY() != 7 {
		t.Errorf("Ожидалось y = 7, получено %d", cordinate.GetY())
	}
}

func TestGetX(t *testing.T) {
	cordinate := NewCoordinate(6, 1)

	if x := cordinate.GetX(); x != 6 {
		t.Errorf("GetX() = %d, ожидалось 6", x)
	}
}

func TestGetY(t *testing.T) {
	cordinate := NewCoordinate(6, 1)

	if y := cordinate.GetY(); y != 1 {
		t.Errorf("GetY() = %d, ожидалось 1", y)
	}
}
