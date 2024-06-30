package parttwo

import (
	"day12/field"
	"day12/order"
	"day12/partone"
	"day12/parttwo/solution"
)

func SolvePart1Varant(f field.Field, s solution.SolutionKey) int {
	f, ok := f.RestrictFromLeftAndRight(s.LeftDemand, s.RightDemand)
	if !ok {
		return 0
	}
	return partone.NumberOfCombination(f, s.SubOrder)
}

func FindAllPossibleNextkeysMiddle(previous solution.SubProblem, remainingOrder order.Order) []solution.SubProblem {
	// här måste vi definiera hur pusselbitarna ska sättas ihop
	// behöver en typ som inte är solutionkey utan subProblem eller liknande
	// det bör finnas ett enklare sätt att få till detta
	var nextSubproblems []solution.SubProblem = []solution.SubProblem{}
	return nextSubproblems
}

func NumberOfCombination(f field.Field, o order.Order, dept int) int {
	o.Unfold(dept)                                               // hur många gånger det skulle upprepas
	original := generation{solution.NoConnectionDDemand(), o, 1} // det här kanske fungerar
	childrenOfTime := []generation{original}
	for i := 0; i < dept; i++ {
		nextChildren := nextGeneration(childrenOfTime, f)
		childrenOfTime = nextChildren
	}
	sum := 0
	for _, el := range childrenOfTime {
		sum += el.nrSolutions
	}
	return sum
}

func nextGeneration(parents []generation, f field.Field) []generation {
	return []generation{}
}

type generation struct {
	prevRightConnection solution.Connection
	remainingOrder      order.Order
	nrSolutions         int
}
