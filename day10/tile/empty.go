package tile

import "day10/point"

type empty struct {
	pos point.Vec
}

func (e empty) Ways() point.VecSet {
	return point.EmptySet()
}

func (e empty) GetPos() point.Vec {
	return e.pos
}
