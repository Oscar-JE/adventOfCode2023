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

func (s Splitter) Reflect(b beam.Beam) []beam.Beam {

}
