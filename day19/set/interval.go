package set

type Interval struct {
	lower int
	upper int
}

func (i Interval) Cardinality() int {
	return max(i.upper-i.lower, 0)
}

func (i Interval) DemandAbove(limit int) Interval {
	return Interval{i.lower, min(limit, i.upper)}
}

func (i Interval) DemandBelow(limit int) Interval {
	newLowerBond := max(i.lower, limit+1)
	return Interval{newLowerBond, i.upper}
}

func (i Interval) IsEmpty() bool {
	return i.Cardinality() == 0
}

func Intersect(a Interval, b Interval) Interval {
	return Interval{lower: max(a.lower, b.lower), upper: min(a.upper, b.upper)}
}

func (i Interval) contains(other Interval) bool {
	return i.lower <= other.lower && other.upper <= i.upper
}

func (i Interval) containsElement(el int) bool {
	return i.lower <= el && el < i.upper
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
