package ticket

import "testing"

func TestParse(t *testing.T) {
	ints1, ints2 := Parse("Card 1: 1 2 | 3 4 5")
	expected1 := []int{1, 2}
	expected2 := []int{3, 4, 5}
	if !arrEq(ints1, expected1) {
		t.Errorf("faulty first int array")
	}
	if !arrEq(ints2, expected2) {
		t.Errorf("faulty second <int array")
	}
}

func arrEq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	eq := true
	for i := range a {
		eq = eq && a[i] == b[i]
	}
	return eq
}
