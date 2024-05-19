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
