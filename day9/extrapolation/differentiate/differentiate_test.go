package differentiate

import "testing"

func TestTrivialCase(t *testing.T) {
	input := []int{}
	res := Diff(input)
	if len(res) != 0 {
		t.Errorf("an empty input should return empty")
	}
}

func TestSimpleCase(t *testing.T) {
	input := []int{1, 1}
	res := Diff(input)
	if len(res) != 1 && res[0] != 0 {
		t.Errorf("Do you not know simple substraction?")
	}
}
