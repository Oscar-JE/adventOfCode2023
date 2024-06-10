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
	nrSolutions := recursivePart(f, o, repitition, 0)
	return nrSolutions
}

func recursivePart(f field.Field, o order.Order, deept int, leftTaken int) int { // den innan biten måste in här
	var parts []solution.SolutionKey
	if deept != 0 {
		parts = findPossibleNextPuzzleMiddleParts(f, o, leftTaken)
	} else {
		parts = findPossibleNextPuzzleMiddleLast(f, o, leftTaken)
	}
	sum := 0
	for _, sol := range parts {
		// måste räkna ut en suborder här
		// logik för att räkna ut vad nästas leftdemand från förras right måste räklnas ut någonstans
		sum += Part1SubProblem(f, sol) * recursivePart(f, o, deept-1, sol.RightDemand) // här bör det ske tekenskillnader
	}
	return sum
}

func findPossibleNextPuzzleMiddleLast(f field.Field, o order.Order, leftTaken int) []solution.SolutionKey {
	unfiltered := findPossibleNextPuzzleMiddleParts(f, o, leftTaken)
	filtered := []solution.SolutionKey{}
	for _, sol := range unfiltered {
		if sol.RightDemand == -1 {
			filtered = append(filtered, sol)
		}
	}
	return filtered
}

func findPossibleNextPuzzleMiddleParts(f field.Field, o order.Order, leftTaken int) []solution.SolutionKey {
	solutions := []solution.SolutionKey{}
	solutions = append(solutions, solution.Init(leftTaken, order.Init([]int{}), -1))
	lastSolution := solutions[len(solutions)-1]
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
