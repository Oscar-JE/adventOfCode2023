package extrapolation

import (
	"testing"
)

func TestTrivialCase(t *testing.T) {
	intput := []int{}
	res := Extrapolate(intput)
	if res != 0 {
		t.Errorf("missed the trivial case")
	}
}

func TestSimpleCase(t *testing.T) {
	input := []int{1, 1}
	res := Extrapolate(input)
	if res != 1 {
		t.Errorf("Who could have guessed 1 should follow from only ones")
	}
}

func TestDeptOne(t *testing.T) {
	input := []int{0, 1, 2}
	res := Extrapolate(input)
	if res != 3 {
		t.Errorf("1,2,3 due")
	}
}

func TestDepthTwo(t *testing.T) { // måste kolla med instructionerna
	// 0  0  1  3  6
	//   0  1  2  3
	//     1  1  1
	// 		 0  0
	input := []int{0, 0, 1, 3}
	res := Extrapolate(input)
	if res != 6 {
		t.Errorf("See the comment in test depth two, answer should be 6")
	}
}

func TestBackwardsExample(t *testing.T) {
	input := []int{10, 13, 16, 21, 30, 45}
	res := Backwards(input)
	if res != 5 {
		t.Errorf("Examine backwards extrapolation more carefully")
	}
}
