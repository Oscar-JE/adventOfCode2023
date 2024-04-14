package tile

import "day10/point"

type empty struct {
}

func (e empty) moves() point.VecSet {
	return point.EmptySet()
}
