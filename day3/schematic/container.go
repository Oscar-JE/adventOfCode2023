package schematic

type NumContainer struct {
	numbers []NumberAndPosition
}

func (nr *NumContainer) add(numPos NumberAndPosition) {
	if !nr.contains(numPos) {
		nr.numbers = append(nr.numbers, numPos)
	}
}
func (nr NumContainer) contains(numPos NumberAndPosition) bool {
	contains := false
	for _, numPosContained := range nr.numbers {
		contains = contains || numPosContained.position == numPos.position
	}
	return contains
}

func (nr NumContainer) getSlize() []int {
	ret := []int{}
	for _, numPos := range nr.numbers {
		ret = append(ret, numPos.number)
	}
	return ret
}

type NumberAndPosition struct {
	number   int
	position Position
}

type Position struct {
	row int
	col int
}

func (p Position) Add(other Position) Position {
	return Position{p.row + other.row, p.col + other.col}
}
