package tile

import "day10/point"

type start struct {
}

func (e start) moves() point.VecSet {
	return point.InitVecSet([]point.Vec{point.Init(1, 0), point.Init(-1, 0), point.Init(0, 1), point.Init(0, -1)})
}
