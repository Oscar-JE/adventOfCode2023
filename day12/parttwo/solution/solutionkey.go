package solution

import (
	"day12/order"
	"fmt"
)

type SolutionKey struct {
	LeftDemand  int
	RightDemand int
	SubOrder    order.Order
}

func (k SolutionKey) String() string {
	return fmt.Sprint(k.LeftDemand, k.RightDemand, k.SubOrder)
}

func Init(leftTaken int, order order.Order, rightTake int) SolutionKey {
	return SolutionKey{leftTaken, rightTake, order}
}

func (s SolutionKey) LeastSpotsOccupied() int {
	var left int
	if s.LeftDemand == -1 {
		left = 0
	} else {
		left = s.LeftDemand + 1
	}
	var right int
	if s.RightDemand == -1 {
		right = 0
	} else {
		right = s.RightDemand
	}
	o := s.SubOrder
	lsum := 0
	for o.Len() > 0 {
		lsum += o.Pop() + 1
	}
	return left + lsum + right
}
