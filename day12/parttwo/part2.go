package parttwo

import (
	"day12/field"
	"day12/order"
	"day12/partone"
	"day12/parttwo/solution"
)

func Part1SubProblem(f field.Field, s solution.SolutionKey) int {
	restriktedFeild, ok := f.RestrictFromLeftAndRight(s.LeftDemand, s.RightDemand)
	if !ok {
		return 0
	}
	return partone.NumberOfCombination(restriktedFeild, s.SubOrder)
}

func Solve(f field.Field, o order.Order, repitition int) int {
	o.Unfold(repitition)
	nrSolutions := recursivePart(f, o, repitition)
	return 5 * len(nextParts)
}

func recursivePart(f field.Field, o order.Order, deept int) int {

}

func findPossibleNextPuzzleMiddleParts(f field.Field, o order.Order, leftTaken int) []solution.SolutionKey {
	// kan göra oss av med f som en variabel och enbart gå på längden
	solutions := []solution.SolutionKey{}
	solutions = append(solutions, solution.Init(leftTaken, order.Init([]int{}), -1))

	lastSolution := solutions[len(solutions)-1] // ja men det börjar kännas bra
	for lastSolution.LeastSpotsOccupied() < f.Len() {
		slidingBlok := o.Pop()
		solutions = append(solutions, solutionsWithAddedBlock(lastSolution, slidingBlok)...)
		lastSolution = solutions[len(solutions)-1]
	}
	return solutions
}

func solutionsWithAddedBlock(lastSolution solution.SolutionKey, slidingBlok int) []solution.SolutionKey {
	solutions := []solution.SolutionKey{}
	for i := range slidingBlok {
		sol := lastSolution
		sol.RightDemand = i
		solutions = append(solutions, sol)
	}
	solutions = append(solutions, solution.Init(lastSolution.LeftDemand, lastSolution.SubOrder.Append(slidingBlok), -1))
	return solutions
}
