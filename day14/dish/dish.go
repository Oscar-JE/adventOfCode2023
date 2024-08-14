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

func (d *Dish) TiltSouth() {
	nrColumns := d.layout.GetCols()
	for j := range nrColumns {
		column := d.layout.GetColumn(j)
		tilted := fallDown(column)
		d.layout.SetColumn(tilted, j)
	}
}

func (d *Dish) TiltWest() {
	nrRows := d.layout.GetRows()
	for i := range nrRows {
		row := d.layout.GetRow(i)
		tilted := fallUpp(row)
		d.layout.SetRow(tilted, i)
	}
}

func (d *Dish) TiltEast() {
	nrRows := d.layout.GetRows()
	for i := range nrRows {
		row := d.layout.GetRow(i)
		tilted := fallDown(row)
		d.layout.SetRow(tilted, i)
	}
}

func (d *Dish) Cycle() {
	d.TiltNorth()
	d.TiltWest()
	d.TiltSouth()
	d.TiltEast()
}

func (d Dish) CycleForever() []matrix.Matrix {
	repetitions := 1001 // 1000000000
	history := []matrix.Matrix{}
	history = append(history, d.layout.DeepCopy())
	for i := 0; i < repetitions; i++ {
		d.Cycle()
		if seen(history, d.layout) {
			history = append(history, d.layout)
			break
		}
		history = append(history, d.layout.DeepCopy())
	}
	return history
}

func seen(history []matrix.Matrix, mat matrix.Matrix) bool {
	for i := range history {
		if history[i].Eq(mat) {
			return true
		}
	}
	return false
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

func (d Dish) Calculate2() int {
	history := d.CycleForever()
	joningIndex := indexOf(history, history[len(history)-1])
	aktiveLen := len(history) - 1
	loopIndex := (1000000000 - joningIndex) % (aktiveLen - joningIndex)
	d.layout = history[loopIndex+joningIndex]
	return d.Score()
}

func (d Dish) String() string {
	return d.layout.String()
}

func indexOf(hist []matrix.Matrix, mat matrix.Matrix) int {
	for i := range hist {
		if hist[i].Eq(mat) {
			return i
		}
	}
	return len(hist)
}
