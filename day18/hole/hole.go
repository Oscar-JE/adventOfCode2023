package hole

import (
	"day16/vec"
	"fmt"
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
	return fmt.Sprintln(minRow + maxRow)
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
