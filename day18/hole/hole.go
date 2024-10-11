package hole

import (
	"day16/vec"
	"day18/segment"
	"fmt"
)

type Hole struct { // bör kunna skriva denna som en serie av segment
	boarder []segment.Segment
}

func Init(boarder []segment.Segment) Hole {
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
	minRow := h.boarder[0].MinRow()
	for i := 1; i < len(h.boarder); i++ {
		row := h.boarder[i].MinRow()
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
	maxRow := h.boarder[0].MaxRow()
	for i := 1; i < len(h.boarder); i++ {
		row := h.boarder[i].MaxRow()
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
	minCol := h.boarder[0].MinCol()
	for i := 1; i < len(h.boarder); i++ {
		col := h.boarder[i].MinCol()
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
	maxCol := h.boarder[0].MaxCol()
	for i := 1; i < len(h.boarder); i++ {
		col := h.boarder[i].MaxCol()
		if maxCol < col {
			maxCol = col
		}
	}
	return maxCol
}

func (h Hole) has(x int, y int) bool {
	point := vec.Init(x, y)
	for _, segment := range h.boarder {
		if segment.In(point) {
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

func (h Hole) findTopSegment() segment.Segment {
	topRow := h.findMinRow()
	for i := 0; i < len(h.boarder); i++ {
		seg := h.boarder[i]
		if seg.MinRow() == topRow && seg.IsHorizontal() {
			return seg
		}
	}
	panic("No top segment where found")
}

func (h Hole) findStartingPoints() []vec.Vec2d {
	topSeg := h.findTopSegment()
	travelDirection := topSeg.GetDirection()
	inwards := vec.Init(1, 0)
	nrRotations := rotationsToEquality(travelDirection, inwards)
	return h.findOneStepInwards(nrRotations)
}

func (h Hole) NrOfInteriorPoints() int { // generationer
	visited := h.findStartingPoints()
	queue := visited
	directions := []vec.Vec2d{vec.Init(1, 0), vec.Init(-1, 0), vec.Init(0, 1), vec.Init(0, -1)}
	for len(queue) > 0 { // bör kunna skriva som en generations cykel istället
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

func (h Hole) NrOfInteriorPoints2() int {
	sum := 0
	directions := []vec.Vec2d{vec.Init(1, 0), vec.Init(-1, 0), vec.Init(0, 1), vec.Init(0, -1)}
	previousGen := []vec.Vec2d{}
	nextGeneration := []vec.Vec2d{}
	currentGen := h.findStartingPoints()
	nrGenerations := 0
	for len(currentGen) != 0 && nrGenerations < 0 {
		sum += len(currentGen)
		for _, point := range currentGen {
			for _, direction := range directions {
				nextGecCandidate := vec.Add(point, direction)
				if !listHas(currentGen, nextGecCandidate) && !listHas(previousGen, nextGecCandidate) && !listHas(nextGeneration, nextGecCandidate) && !h.inBoarder(nextGecCandidate) {
					nextGeneration = append(nextGeneration, nextGecCandidate)
				}
			}
		}
		previousGen = currentGen
		currentGen = nextGeneration
		nextGeneration = []vec.Vec2d{}
		nrGenerations++
	}
	return sum
}

func (h Hole) NrExcavated() int {
	return h.NrOfInteriorPoints2() + h.lengthOfBoarder()
}

func (h Hole) lengthOfBoarder() int {
	sum := 0
	for _, seg := range h.boarder {
		sum += seg.NrOfPoints()
	}
	return sum
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
	for _, seg := range h.boarder {
		inwards := seg.GetDirection().Rotate90DegresUpMultiple(rotation)
		interiorCandidates := seg.Translate(inwards)
		for _, point := range interiorCandidates {
			if !listHas(oneStepIn, point) && !h.inBoarder(point) {
				oneStepIn = append(oneStepIn, point)
			}
		}

	}
	return oneStepIn
}

func (h Hole) inBoarder(v vec.Vec2d) bool {
	for _, seg := range h.boarder {
		if seg.In(v) {
			return true
		}
	}
	return false
}

func listHas(list []vec.Vec2d, point vec.Vec2d) bool {
	for _, p := range list {
		if p == point {
			return true
		}
	}
	return false
}
