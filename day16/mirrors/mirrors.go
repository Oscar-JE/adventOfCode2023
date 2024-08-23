package mirrors

import (
	"day16/beam"
	"day16/matrix"
)

type Mirrors struct {
	mat   matrix.Matrix[Reflector]
	beams []beam.Beam
}

type Reflector interface {
	Reflect(beam.Beam) []beam.Beam
}

func (m *Mirrors) NextState() {
	m.moveBeams()
	m.removeBeamsOutOfBound()
	m.reflectBeams()
}

func (m *Mirrors) moveBeams() {
	for _, beam := range m.beams {
		beam.Travel()
	}
}

func (m *Mirrors) removeBeamsOutOfBound() {
	beamsInside := []beam.Beam{}
	for _, beam := range m.beams {
		position := beam.GetPosition()
		if m.mat.Inside(position.GetX(), position.GetY()) {
			beamsInside = append(beamsInside, beam)
		}
	}
	m.beams = beamsInside
}

func (m *Mirrors) reflectBeams() {
	reflectedBeams := []beam.Beam{}
	for _, beam := range m.beams {
		reflector := m.mat.Get(beam.GetDirection().GetX(), beam.GetDirection().GetY())
		reflectedBeams = append(reflectedBeams, reflector.Reflect(beam)...)
	}
	m.beams = reflectedBeams
}
