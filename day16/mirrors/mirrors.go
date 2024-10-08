package mirrors

import (
	"day16/beam"
	beamcacher "day16/beamcach"
	"day16/energyreport"
	"day16/matrix"
	"day16/vec"
)

type Reflector interface {
	Reflect(vec.Vec2d) []vec.Vec2d
	String() string
}

type NrEnergized interface {
	ReportPosition(vec.Vec2d)
	TotalNrEnergized() int
}

type Mirrors struct {
	mat      matrix.Matrix[Reflector]
	beams    []beam.Beam
	reporter NrEnergized
	cach     beamcacher.BeamCacher
}

func Init(m matrix.Matrix[Reflector], starterBeam beam.Beam) Mirrors {
	rows := m.GetRows()
	cols := m.GetCols()
	reporter := energyreport.Init(rows, cols) // kommer denna bli gb:ad ??
	beams := []beam.Beam{starterBeam}
	cach := beamcacher.Init(rows, cols)
	return Mirrors{mat: m, beams: beams, reporter: &reporter, cach: cach}
}

func (m *Mirrors) NextState() {
	m.moveBeams()
	m.removeBeamsOutOfBound()
	m.reportBeams()
	m.reflectBeams()
}

func (m *Mirrors) moveBeams() { // denna fungerar inte som jag förväntade
	for i := range m.beams {
		m.beams[i].Travel()
	}
}

func (m *Mirrors) reportBeams() {
	for _, beam := range m.beams {
		m.reporter.ReportPosition(beam.GetPosition())
		m.cach.See(beam)
	}
}

func (m *Mirrors) removeBeamsOutOfBound() {
	beamsInside := []beam.Beam{}
	for _, beam := range m.beams {
		position := beam.GetPosition()
		if m.mat.Inside(position.GetX(), position.GetY()) && (!m.cach.HasSeen(beam)) {
			beamsInside = append(beamsInside, beam)
		}
	}
	m.beams = beamsInside
}

func (m *Mirrors) reflectBeams() {
	reflectedBeams := []beam.Beam{}
	for _, beam := range m.beams {
		reflectedBeams = append(reflectedBeams, m.reflect(beam)...)
	}
	m.beams = reflectedBeams
}

func (m Mirrors) reflect(b beam.Beam) []beam.Beam {
	reflector := m.mat.Get(b.GetPosition().GetX(), b.GetPosition().GetY())
	directions := reflector.Reflect(b.GetDirection())
	beams := []beam.Beam{}
	for _, v := range directions {
		beams = append(beams, beam.Init(b.GetPosition(), v))
	}
	return beams
}

func (m Mirrors) NrOfEnergised() int {
	for i := 0; i < m.mat.GetRows()*m.mat.GetCols(); i++ {
		m.NextState()
		if len(m.beams) == 0 {
			break
		}
	}
	return m.reporter.TotalNrEnergized()
}
