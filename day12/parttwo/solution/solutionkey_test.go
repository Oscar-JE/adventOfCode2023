package solution

import (
	"day12/order"
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	o := order.Init([]int{1, 2, 3})
	key := SolutionKey{0, 1, o}
	fmt.Println(key)
}

func TestSpotsOccupiedEmpty(t *testing.T) {
	key := SolutionKey{-1, -1, order.Init([]int{})}
	res := key.LeastSpotsOccupied()
	if res != 0 {
		t.Errorf("was %d should be 5", res)
	}
}

func TestSpotsOccupiedNibbling(t *testing.T) {
	key := SolutionKey{1, 1, order.Init([]int{})}
	res := key.LeastSpotsOccupied()
	if res != 3 {
		t.Errorf("four should be taken")
	}
}

func TestSpotsOccupiedFullHouse(t *testing.T) {
	key := SolutionKey{1, 1, order.Init([]int{1})}
	res := key.LeastSpotsOccupied()
	if res != 5 {
		t.Errorf("5 should be taken was %d", res)
	}
}
