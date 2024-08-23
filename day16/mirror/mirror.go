package mirror

import (
	"day16/beam"
	"day16/vec"
)

type Empty struct {
}

func (e Empty) Reflect(b beam.Beam) []beam.Beam {
	return []beam.Beam{b}
}

type Splitter struct {
	splitDirection vec.Vec2d
}

func (s Splitter) Reflect(b beam.Beam) []beam.Beam { // får läsa på om det är rätt beteende
	if b.GetDirection().Orthogonal(s.splitDirection) {
		position := b.GetPosition()
		dir1 := s.splitDirection
		dir2 := s.splitDirection.Flip()
		return []beam.Beam{beam.Init(position, dir1), beam.Init(position, dir2)}
	}
	return []beam.Beam{}
}

type Mirror struct {
	normal vec.Vec2d
}
