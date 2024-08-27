package energyreport

import (
	"day16/matrix"
	"day16/vec"
)

type Duality struct {
	m matrix.Matrix[bool]
}

func Init(rows int, cols int) Duality {
	data := []bool{}
	for i := 0; i < rows*cols; i++ {
		data = append(data, false)
	}
	mat := matrix.Init(data, rows, cols)
	return Duality{m: mat}
}

func (d *Duality) ReportPosition(v vec.Vec2d) {
	d.m.Set(v.GetX(), v.GetY(), true)
}

func (d *Duality) TotalNrEnergized() int {
	sum := 0
	for _, energised := range d.m.GetData() {
		if energised {
			sum++
		}
	}
	return sum
}
