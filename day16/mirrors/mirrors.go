package mirrors

import (
	"day16/beam"
	"day16/matrix"
)

type Mirrors struct {
	m     matrix.Matrix[Reflector]
	beams []beam.Beam
}

type Reflector interface {
	Reflect(beam.Beam) []beam.Beam
}
