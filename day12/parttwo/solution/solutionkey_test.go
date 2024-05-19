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
