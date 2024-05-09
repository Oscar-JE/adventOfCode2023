package fields

import (
	"day12/field"
	"day12/order"
	"day12/partone"
	"testing"
)

func TestNrOfSolutions(t *testing.T) {
	p := PiceOfPuzzle{left: Connection{0, false, 0}, right: Connection{0, false, 0}}
	f := field.Init("###.???")
	o := order.Init([]int{3, 1})
	res := NrOfSolutions(f, o, p)
	expected := partone.NumberOfCombination(f, o)
	if res != expected {
		t.Errorf("Nr of  solutions not consistent between part1 and trivial case for flat/flat pice in part2")
	}
}
