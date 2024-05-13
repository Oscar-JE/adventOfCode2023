package fields

import (
	"day12/field"
	"day12/order"
	"day12/partone"
	"errors"
)

type PiceOfPuzzle struct {
	left  Connection
	right Connection
}

func PossiblePices(f field.Field, o order.Order) []PiceOfPuzzle {
	// kommer den här att vara vettig?
	return []PiceOfPuzzle{}
}

func NrOfSolutions(f field.Field, o order.Order, p PiceOfPuzzle) int {
	// här börjar det bli intressant. Det är ett garanterat problem vi måste lösa
	// strategi. Manipulera field och order utifrån PiceOfPuzzle för att åttergå till problem av typ1
	// börjar med att skriva William test för det fallet vi redan löst

	//manipulate left
	// det blir inte snyggt vi definierar istället  en positiv riktning
	return partone.NumberOfCombination(f, o)
}

func modifyOrder(o order.Order, p PiceOfPuzzle) (order.Order, error) { // hur testas denna korrekt
	nrOfBlocksLeft := o.Len() - p.left.nrOfMoved + p.right.nrOfMoved
	if nrOfBlocksLeft < 0 {
		return o, errors.New("impossible order manipulation")
	}
	res := o.RolingRsize(p.left.nrOfMoved, p.right.nrOfMoved)
	return res, nil
}
