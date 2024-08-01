package matrix

type Matrix struct {
	elements []int
	rows     int
	cols     int
	rowOrder bool
}

func Init() Matrix {
	return Matrix{elements: []int{1, 2, 3}, rows: 1, cols: 3, rowOrder: true}
}

func InitZero(rows int, cols int) Matrix {
	nrOfElements := rows * cols
	els := []int{}
	i := 0
	for i < nrOfElements {
		els = append(els, 0)
		i++
	}
	return Matrix{elements: els, rows: rows, cols: cols, rowOrder: true}
}

func InitRowOrder(rows int, cols int, values []int) Matrix {
	if rows*cols != len(values) {
		panic("Illegal matrix initialization")
	}
	return Matrix{elements: values, rows: rows, cols: cols, rowOrder: true}
}

func (m Matrix) indexMapping(row int, col int) int {
	if m.rowOrder {
		return row*m.cols + col
	} else {
		return col*m.rows + row
	}
}

func (m Matrix) Get(row int, col int) int {
	return m.elements[m.indexMapping(row, col)]
}

func (m *Matrix) Set(row int, col int, value int) {
	m.elements[m.indexMapping(row, col)] = value
}

func (m *Matrix) Transpose() {
	cols := m.cols
	m.cols = m.rows
	m.rows = cols
	m.rowOrder = !m.rowOrder
}

func (m Matrix) GetRow(row int) []int {
	rowList := []int{}
	for i := range m.cols {
		rowList = append(rowList, m.Get(row, i))
	}
	return rowList
}

func (m Matrix) GetNrRows() int {
	return m.rows
}
