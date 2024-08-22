package matrix

type Matrix[T any] struct {
	data []T
	rows int
	cols int
}

func Init[T any](data []T, rows int, cols int) Matrix[T] {
	return Matrix[T]{data: data, rows: rows, cols: cols}
}

func (m Matrix[T]) Get(row int, col int) T {
	return m.data[m.index(row, col)]
}

func (m Matrix[T]) index(row int, col int) int {
	return m.rows*m.cols + col
}

func (m *Matrix[T]) Set(row int, col int, value T) {
	m.data[m.index(row, col)] = value
}

func (m Matrix[T]) Inside(row int, col int) bool {
	return (0 <= row && row < m.rows) && (0 <= col && col < m.cols)
}
