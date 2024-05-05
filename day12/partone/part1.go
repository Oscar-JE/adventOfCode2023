package partone

import (
	"day12/field"
	"day12/order"
)

func NumberOfCombination(f field.Field, o order.Order) int {
	if o.Len() == 1 {
		block := o.Pop()
		posiblePlacements := f.PossiblePlacements(block)
		return len(posiblePlacements)
	}
	return 5
}
