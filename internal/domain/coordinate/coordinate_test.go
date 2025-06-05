package coordinate

import "testing"

func TestNewCoordinate(t *testing.T) {
	c := NewCoordinate(3, 7)

	if c.GetX() != 3 {
		t.Errorf("Ожидалось х = 3, получено %d", c.GetX())
	}

	if c.GetY() != 7 {
		t.Errorf("Ожидалось y = 7, получено %d", c.GetY())
	}
}

func TestGetX(t *testing.T) {
	c := NewCoordinate(6, 1)

	if x := c.GetX(); x != 6 {
		t.Errorf("GetX() = %d, ожидалось 6", x)
	}
}

func TestGetY(t *testing.T) {
	c := NewCoordinate(6, 1)

	if y := c.GetY(); y != 1 {
		t.Errorf("GetY() = %d, ожидалось 1", y)
	}
}
