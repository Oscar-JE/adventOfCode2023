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

func (u Univers) Expand() Univers {
	copy := Init(u.galaxyLocations) // blir djup kopia
	copy.galaxyLocations = append(copy.galaxyLocations, linearalgebra.InitVec(0, 0))
	nrRows, nrCols := u.findMaxRowsAndCols()
	rowExpander := linearalgebra.InitVec(1, 0)
	colExpander := linearalgebra.InitVec(0, 1)
	return copy
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
		maxX = max(maxX, u.galaxyLocations[i].GetY())
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
