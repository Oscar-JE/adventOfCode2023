package parttwo

import (
	"day12/field"
	"day12/order"
	"day12/partone"
)

func NumberOfCombination(f field.Field, o order.Order) int {
	f.Unfold(4)
	o.Unfold(4)
	return partone.NumberOfCombination(f, o)
}
