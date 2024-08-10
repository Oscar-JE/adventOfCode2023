package matrix

import (
	"day14/dish/pice"
	"fmt"
)

type Matrix struct {
	pices []pice.Pice
	rows  int
	cols  int
}

func Init(pices []pice.Pice, rows int, cols int) Matrix {
	if rows*cols != len(pices) {
		panic("Rows and Cols not compatabel with number of elements in Matrix")
	}
	return Matrix{pices: pices, rows: rows, cols: cols}
}

func (m Matrix) Get(row int, col int) pice.Pice {
	return m.pices[row*m.cols+col]
}

func (m *Matrix) Set(row int, col int, p pice.Pice) {
	m.pices[row*m.cols+col] = p
}

func (m Matrix) GetCols() int {
	return m.cols
}

func (m Matrix) GetRows() int {
	return m.rows
}

func (m Matrix) GetColumn(cIndex int) []pice.Pice {
	column := []pice.Pice{}
	for i := range m.rows {
		column = append(column, m.Get(i, cIndex))
	}
	return column
}

func (m Matrix) GetRow(rIndex int) []pice.Pice {
	row := []pice.Pice{}
	for j := range m.cols {
		row = append(row, m.Get(rIndex, j))
	}
	return row
}

func (m *Matrix) SetColumn(column []pice.Pice, cIndex int) {
	for row := range m.rows {
		m.Set(row, cIndex, column[row])
	}
}

func (m *Matrix) SetRow(row []pice.Pice, rIndex int) {
	for col := range m.cols {
		m.Set(rIndex, col, row[col])
	}
}

func (m Matrix) String() string {
	rep := ""
	for rowIndex := range m.rows {
		rep += fmt.Sprintln(m.GetRow(rowIndex))
	}
	return rep
}

func Invert(pices []pice.Pice) []pice.Pice {
	inverted := []pice.Pice{}
	for i := len(pices) - 1; i >= 0; i-- {
		inverted = append(inverted, pices[i])
	}
	return inverted
}

func (m Matrix) Eq(other Matrix) bool {
	if m.rows != other.rows || m.cols != other.cols {
		return false
	}
	for i, el := range m.pices {
		if el != other.pices[i] {
			return false
		}
	}
	return true
}

func (m Matrix) DeepCopy() Matrix {
	newInternal := []pice.Pice{}
	for _, el := range m.pices {
		newInternal = append(newInternal, el)
	}
	return Init(newInternal, m.rows, m.cols)
}
