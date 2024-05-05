package partone

import (
	"day12/field"
	"day12/order"
)

func NumberOfCombination(f field.Field, o order.Order) int {
	if o.Len() == 0 {
		if f.HasDamaged() {
			return 0
		} else {
			return 1
		}
	}
	sum := 0
	block := o.Pop()
	posiblePlacements := f.PossiblePlacements(block)
	for _, placement := range posiblePlacements {
		subField := f.SubField(placement + block + 1)
		sum += NumberOfCombination(subField, o)
	}
	return sum
}
