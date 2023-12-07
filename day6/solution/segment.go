package solution

import "math"

type Segment struct {
	lower float64
	upper float64
}

func (s Segment) NrOfIntsInside() int {
	lowestInt := int(math.Floor(s.lower+1) + 0.1)
	largestint := int(math.Ceil(s.upper-1) + 0.1)
	return largestint - lowestInt + 1
}
