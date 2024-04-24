package universe

import (
	"day11/linearalgebra"
	"fmt"
	"testing"
)

func TestStringRep(t *testing.T) {
	u := Init([]linearalgebra.Vec{linearalgebra.InitVec(3, 3)})
	fmt.Println(u)
}

func TestStringRepFromInstructions(t *testing.T) {
	positions := []linearalgebra.Vec{linearalgebra.InitVec(0, 3),
		linearalgebra.InitVec(1, 7),
		linearalgebra.InitVec(2, 0),
		linearalgebra.InitVec(4, 6),
		linearalgebra.InitVec(5, 1),
		linearalgebra.InitVec(6, 9),
		linearalgebra.InitVec(8, 7),
		linearalgebra.InitVec(9, 0),
		linearalgebra.InitVec(9, 4)}
	u := Init(positions)
	fmt.Println(u)
}

func TestExpand(t *testing.T) {
	u := Init([]linearalgebra.Vec{linearalgebra.InitVec(1, 1)})
	expanded := u.Expand()
	fmt.Println(u)
	fmt.Println(expanded)
}
