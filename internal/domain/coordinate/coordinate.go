package coordinate

type Coordinate struct {
	xcoordinate int
	ycoordinate int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{
		xcoordinate: x,
		ycoordinate: y,
	}
}

func (c *Coordinate) GetX() int { return c.xcoordinate }

func (c *Coordinate) GetY() int { return c.ycoordinate }
