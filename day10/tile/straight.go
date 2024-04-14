package tile

import "day10/point"

type straigth struct {
	standing bool
}

func (s straigth) moves() point.VecSet {
	set := point.EmptySet()
	if s.standing {
		set.Add(point.Init(1, 0))
		set.Add(point.Init(-1, 0))
	} else {
		set.Add(point.Init(0, 1))
		set.Add(point.Init(0, -1))
	}
	return set
}
