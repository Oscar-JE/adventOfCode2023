package matrix

import "fmt"

type Matrix[T fmt.Stringer] struct {
	data []T
	rows int
	cols int
}

func Init[T fmt.Stringer](data []T, rows int, cols int) Matrix[T] {
	return Matrix[T]{data: data, rows: rows, cols: cols}
}

func (m Matrix[T]) Get(row int, col int) T {
	return m.data[m.index(row, col)]
}

func (m Matrix[T]) GetRows() int {
	return m.rows
}

func (m Matrix[T]) GetCols() int {
	return m.cols
}

func (m Matrix[T]) index(row int, col int) int {
	return m.cols*row + col
}

func (m *Matrix[T]) Set(row int, col int, value T) {
	m.data[m.index(row, col)] = value
}

func (m Matrix[T]) Inside(row int, col int) bool {
	return (0 <= row && row < m.rows) && (0 <= col && col < m.cols)
}

func (m Matrix[T]) GetData() []T {
	return m.data
}

func (m Matrix[T]) String() string {
	representation := ""
	for row := 0; row < m.rows; row++ {
		rowRep := ""
		for col := 0; col < m.cols; col++ {
			rowRep += m.Get(row, col).String()
		}
		representation += rowRep + "\r\n"
	}
	return representation
}
