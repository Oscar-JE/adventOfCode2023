package dish

import (
	"day14/dish/matrix"
	"day14/dish/pice"
	"strings"
)

type Dish struct {
	layout matrix.Matrix
}

func Init(l matrix.Matrix) Dish {
	return Dish{layout: l}
}

func InitFromText(rep string) Dish {
	splitted := strings.Split(rep, "\n")
	rows := len(splitted)
	firstRunes := []rune(splitted[0])
	cols := len(firstRunes)
	stripped := strings.Replace(rep, "\n", "", -1)
	pices := []pice.Pice{}
	for _, r := range stripped {
		pices = append(pices, pice.InitFromRune(r))
	}
	return Init(matrix.Init(pices, rows, cols))
}

func (d *Dish) TiltNorth() {
	nrColumns := d.layout.GetCols()
	for j := range nrColumns {
		column := d.layout.GetColumn(j)
		tilted := fallUpp(column)
		d.layout.SetColumn(tilted, j)
	}
}

func (d Dish) Score() int {
	score := 0
	nrColumns := d.layout.GetCols()
	for j := range nrColumns {
		column := d.layout.GetColumn(j)
		score += arrScore(column)
	}
	return score
}

func (d Dish) Calculate() int {
	d.TiltNorth()
	return d.Score()
}
