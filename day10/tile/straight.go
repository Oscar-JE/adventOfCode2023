package tile

import "day10/point"

type straigth struct {
	standing bool
	position point.Vec
}

func (s straigth) Ways() point.VecSet {
	set := point.EmptySet()
	if s.standing {
		set.Add(s.position.Add(point.Init(1, 0)))
		set.Add(s.position.Add(point.Init(-1, 0)))
	} else {
		set.Add(s.position.Add(point.Init(0, 1)))
		set.Add(s.position.Add(point.Init(0, -1)))
	}
	return set
}

func (s straigth) GetPos() point.Vec {
	return s.position
}
