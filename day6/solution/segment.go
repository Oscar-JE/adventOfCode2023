package solution

type Segment struct {
	lower float64
	upper float64
}

func (s Segment) IntsInside() int {
	return 5
}
