package hole

import (
	"day16/vec"
	"day18/integration"
)

type Hole struct {
	boarder []vec.Vec2d
}

func Init(boarder []vec.Vec2d) Hole {
	return Hole{boarder: boarder}
}

func (h Hole) String() string {
	minRow := h.findMinRow()
	maxRow := h.findMaxRow()
	minCol := h.findMinCol()
	maxCol := h.findMaxCol()
	rep := ""
	for i := minRow; i <= maxRow; i++ {
		for j := minCol; j <= maxCol; j++ {
			if h.has(i, j) {
				rep += "#"
			} else {
				rep += "."
			}
		}
		rep += "\r\n"
	}
	return rep
}

func (h Hole) findMinRow() int {
	if len(h.boarder) == 0 {
		panic("wtf do you expect me to do. There is no least row")
	}
	minRow := h.boarder[0].GetX()
	for i := 1; i < len(h.boarder); i++ {
		row := h.boarder[i].GetX()
		if minRow > row {
			minRow = row
		}
	}
	return minRow
}

func (h Hole) findMaxRow() int {
	if len(h.boarder) == 0 {
		panic("wtf do you expect me to do. There is no largest row")
	}
	maxRow := h.boarder[0].GetX()
	for i := 1; i < len(h.boarder); i++ {
		row := h.boarder[i].GetX()
		if maxRow < row {
			maxRow = row
		}
	}
	return maxRow
}

func (h Hole) findMinCol() int {
	if len(h.boarder) == 0 {
		panic("wtf do you expect me to do. There is no least row")
	}
	minCol := h.boarder[0].GetY()
	for i := 1; i < len(h.boarder); i++ {
		col := h.boarder[i].GetY()
		if minCol > col {
			minCol = col
		}
	}
	return minCol
}

func (h Hole) findMaxCol() int {
	if len(h.boarder) == 0 {
		panic("wtf do you expect me to do. There is no least row")
	}
	maxCol := h.boarder[0].GetY()
	for i := 1; i < len(h.boarder); i++ {
		col := h.boarder[i].GetY()
		if maxCol < col {
			maxCol = col
		}
	}
	return maxCol
}

func (h Hole) has(x int, y int) bool {
	el := vec.Init(x, y)
	for _, point := range h.boarder {
		if el == point {
			return true
		}
	}
	return false
}

func (h Hole) Fill() int {
	minRow := h.findMinRow()
	maxRow := h.findMaxCol()
	sum := 0
	for row := minRow; row <= maxRow; row++ {
		boarderCols := h.findColNumbers(row)
		sum += integration.RowIntegration(boarderCols)
	}
	return sum
}

func (h Hole) findColNumbers(row int) []int {
	collNumbers := []int{}
	for _, point := range h.boarder {
		if point.GetX() == row {
			collNumbers = append(collNumbers, point.GetY())
		}
	}
	return collNumbers
}
