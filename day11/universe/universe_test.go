package universe

import (
	"day11/linearalgebra"
	"fmt"
	"testing"
)

func unexpanded() Univers {
	positions := []linearalgebra.Vec{linearalgebra.InitVec(0, 3),
		linearalgebra.InitVec(1, 7),
		linearalgebra.InitVec(2, 0),
		linearalgebra.InitVec(4, 6),
		linearalgebra.InitVec(5, 1),
		linearalgebra.InitVec(6, 9),
		linearalgebra.InitVec(8, 7),
		linearalgebra.InitVec(9, 0),
		linearalgebra.InitVec(9, 4)}
	return Init(positions)
}

func TestStringRep(t *testing.T) {
	u := Init([]linearalgebra.Vec{linearalgebra.InitVec(3, 3)})
	fmt.Println(u)
}

func TestStringRepFromInstructions(t *testing.T) {
	u := unexpanded()
	fmt.Println(u)
}

func TestExpand(t *testing.T) {
	u := Init([]linearalgebra.Vec{linearalgebra.InitVec(1, 1)})
	expanded := u.Expand(1)
	fmt.Println(u)
	fmt.Println(expanded)
}

func TestExpandedInstructionExample(t *testing.T) {
	u := unexpanded()
	expanded := u.Expand(1)
	fmt.Println(expanded)
}

func TestInstructionShortImput(t *testing.T) {
	u := unexpanded()
	res := u.ExpandAndSumDistance()
	if res != 374 {
		t.Errorf("examine ExtandedAndSumDistance")
	}
}
