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

func NumberOfCombinations2(f field.Field, o order.Order, repetitions int) int {
	o.Unfold(repetitions - 1)
	return positionOfNumbers(f, o, repetitions)
}

func positionOfNumbers(f field.Field, o order.Order, depth int) int {
	possibleOrderConsumptionsAndDemandsOnNext := []FieldOrderAndNrOfSolutions{}
	return len(possibleOrderConsumptionsAndDemandsOnNext)
}

type FieldOrderAndNrOfSolutions struct {
	f           field.Field // behövs detta fält sparas? svar egentligen inte
	o           order.Order
	demand      int // -1 , flat , 0 # middle next need to be '.' , 1 M:# f:#, s:'.' och så vidare
	nrSolutions int // enbart en cashning
}
