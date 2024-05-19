package parttwo

import (
	"day12/field"
)

func SolvePart1Varant(f field.Field, s SolutionKey) int {
	f, ok := f.RestrictFromLeftAndRight(s.LeftDemand, s.RightDemand)
}
