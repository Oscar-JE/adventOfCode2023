package universe

import (
	"day11/linearalgebra"
	"strings"
)

type Univers struct {
	galaxyLocations []linearalgebra.Vec
}

func Init(galaxyLocations []linearalgebra.Vec) Univers {
	return Univers{galaxyLocations: galaxyLocations}
}

func (u Univers) ExpandAndSumDistance() int {
	expanded := u.Expand(1)
	return expanded.SumDistance()
}

func (u Univers) SillyExpandAndSumDistance(expansionRate int) int {
	expanded := u.Expand(expansionRate)
	return expanded.SumDistance()
}

func (expanded Univers) SumDistance() int {
	sum := 0
	for i := range expanded.galaxyLocations {
		for j := i + 1; j < len(expanded.galaxyLocations); j++ {
			sum += linearalgebra.ManhattanDistans(expanded.galaxyLocations[i], expanded.galaxyLocations[j])
		}
	}
	return sum
}

func (u Univers) Expand(multiple int) Univers {
	copy := Init(u.galaxyLocations)
	nrRows, nrCols := u.findMaxRowsAndCols()
	for i := nrRows - 1; i >= 0; i-- {
		if u.emptyrow(i) {
			copy.expandRowsBelow(i, multiple)
		}
	}
	for j := nrCols - 1; j >= 0; j-- {
		if u.emptyCol(j) {
			copy.expandColsTOTheRight(j, multiple)
		}
	}
	return copy
}

func (u Univers) emptyrow(row int) bool {
	for _, pos := range u.galaxyLocations {
		if pos.GetX() == row {
			return false
		}
	}
	return true
}

func (u *Univers) expandRowsBelow(row int, multiple int) {
	howManyToAdd := multiple - 1
	rowExpander := linearalgebra.InitVec(1, 0).Multiply(howManyToAdd)
	for i := range u.galaxyLocations {
		if u.galaxyLocations[i].GetX() > row {
			u.galaxyLocations[i] = linearalgebra.Add(u.galaxyLocations[i], rowExpander)
		}
	}
}

func (u Univers) emptyCol(col int) bool {
	for _, pos := range u.galaxyLocations {
		if pos.GetY() == col {
			return false
		}
	}
	return true
}

func (u *Univers) expandColsTOTheRight(col int, multiple int) {
	howManyToAdd := multiple - 1
	colExpander := linearalgebra.InitVec(0, 1).Multiply(howManyToAdd)
	for i := range u.galaxyLocations {
		if u.galaxyLocations[i].GetY() > col {
			u.galaxyLocations[i] = linearalgebra.Add(u.galaxyLocations[i], colExpander)
		}
	}
}

func (u Univers) String() string {
	maxX, maxY := u.findMaxRowsAndCols()
	universRep := []string{}
	for i := 0; i < maxX; i++ {
		universRep = append(universRep, strings.Repeat(".", maxY))
	}
	for _, vec := range u.galaxyLocations {
		line := []rune(universRep[vec.GetX()])
		line[vec.GetY()] = '#'
		universRep[vec.GetX()] = string(line)
	}
	retString := ""
	for _, el := range universRep {
		retString += el + "\n"
	}
	return retString
}

func (u Univers) findMaxRowsAndCols() (int, int) {
	if len(u.galaxyLocations) == 0 {
		return 0, 0
	}
	maxX := u.galaxyLocations[0].GetX()
	maxY := u.galaxyLocations[0].GetY()
	for i := 1; i < len(u.galaxyLocations); i++ {
		maxX = max(maxX, u.galaxyLocations[i].GetX())
		maxY = max(maxY, u.galaxyLocations[i].GetY())
	}
	return maxX + 1, maxY + 1
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
