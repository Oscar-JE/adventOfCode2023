package set

type Interval struct {
	lower int
	upper int
}

func Intersect(a Interval, b Interval) Interval {
	return Interval{lower: max(a.lower, b.lower), upper: min(a.upper, b.upper)}
}

func (i Interval) contains(other Interval) bool {

}

func (i Interval) Remove(other Interval) []Interval {
	return []Interval{}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
