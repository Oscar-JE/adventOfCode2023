package beamcacher

import (
	"day16/beam"
	"day16/vec"
)

type BeamCacher struct {
	seen []bool
	rows int
	cols int
}

func Init(rows int, cols int) BeamCacher { // fyran för antal riktningar
	seen := []bool{}
	for i := 0; i < rows*cols*4; i++ {
		seen = append(seen, false)
	}
	return BeamCacher{seen: seen, rows: rows, cols: cols}
}

func (c BeamCacher) HasSeen(l beam.Beam) bool {
	i := c.index(l)
	return c.seen[i]
}

func (c *BeamCacher) See(l beam.Beam) {
	c.seen[c.index(l)] = true
}

func (b BeamCacher) positionScore(v vec.Vec2d) int {
	return v.GetX()*b.cols + v.GetY()
}

func directionScore(v vec.Vec2d) int {
	if v.GetX() > 0 {
		return 0
	} else if v.GetX() < 0 {
		return 1
	} else if v.GetY() > 0 {
		return 2
	} else if v.GetY() < 0 {
		return 3
	}
	panic("imposible imput")
}

func (c BeamCacher) index(b beam.Beam) int {
	return c.positionScore(b.GetPosition())*4 + directionScore(b.GetDirection())
}
