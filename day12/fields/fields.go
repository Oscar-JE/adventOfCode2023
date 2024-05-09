package fields

import (
	"day12/field"
	"day12/order"
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
	// börjar med att skriva test för det fallet vi redan löst

	//manipulate left

	return 0
}
