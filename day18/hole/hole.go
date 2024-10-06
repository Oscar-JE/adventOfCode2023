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

func (h Hole) PrintDoubleBoarders() {
	fmt.Println("double digg in following lokations:")
	for i := range h.boarder {
		kontrolElement := h.boarder[i]
		for j := i + 1; j < len(h.boarder); j++ {
			if kontrolElement == h.boarder[j] {
				fmt.Println(kontrolElement)
			}
		}
	}
}

func (h Hole) findIndexOfPairadjacentToTopBoarder() (vec.Vec2d, vec.Vec2d) {
	topRow := h.findMinRow()
	for i := 0; i < len(h.boarder)-1; i++ {
		n1 := h.boarder[i]
		n2 := h.boarder[i+1]
		if n1.GetX() == topRow && n2.GetX() == topRow {
			return n1, n2
		}
	}
	panic("No top boarder neighbors where found")
}

func (h Hole) findStartingPoints() []vec.Vec2d {
	n1, n2 := h.findIndexOfPairadjacentToTopBoarder()
	travelDirection := vec.Subtract(n2, n1) // räkna med hur mycket vi ska vrida
	inwards := vec.Init(1, 0)
	nrRotations := rotationsToEquality(travelDirection, inwards)
	return h.findOneStepInwards(nrRotations)
}

func (h Hole) NrOfInteriorPoints() int {
	visited := h.findStartingPoints()
	queue := visited
	directions := []vec.Vec2d{vec.Init(1, 0), vec.Init(-1, 0), vec.Init(0, 1), vec.Init(0, -1)}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			new := vec.Add(p, dir)
			if !listHas(visited, new) && !h.inBoarder(new) {
				visited = append(visited, new)
				queue = append(queue, new)
			}
		}
	}
	return len(visited)
}

func (h Hole) NrExcavated() int {
	return h.NrOfInteriorPoints() + len(h.boarder)
}

func rotationsToEquality(v1 vec.Vec2d, seeked vec.Vec2d) int {
	nr := 0
	for v1 != seeked {
		v1 = v1.Rotate90DegreesUp()
		if nr >= 4 {
			panic("illegal input: there is no rotation by 90 degrees to equality")
		}
		nr++
	}
	return nr
}

func (h Hole) findOneStepInwards(rotation int) []vec.Vec2d {
	oneStepIn := []vec.Vec2d{}
	for i := 0; i < len(h.boarder)-1; i++ {
		direction := vec.Subtract(h.boarder[i+1], h.boarder[i])
		inwards := direction.Rotate90DegresUpMultiple(rotation)
		interiorCandidate := vec.Add(h.boarder[i+1], inwards)
		if !listHas(oneStepIn, interiorCandidate) && !h.inBoarder(interiorCandidate) {
			oneStepIn = append(oneStepIn, interiorCandidate)
		}
	}
	return oneStepIn
}

func (h Hole) inBoarder(v vec.Vec2d) bool {
	return listHas(h.boarder, v)
}

func listHas(list []vec.Vec2d, point vec.Vec2d) bool {
	for _, el := range list {
		if el == point {
			return true
		}
	}
	return false
}
