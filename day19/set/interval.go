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

func (i Interval) unionOfConnected(other Interval) Interval {
	return Interval{min(i.lower, other.lower), max(i.upper, other.upper)}
}

func (i Interval) isLowerThan(other Interval) bool {
	return i.upper <= other.lower
}

func (i Interval) isLowerThanEl(element int) bool {
	return i.upper <= element
}

func (i Interval) isLargerThanEl(element int) bool {
	return i.lower >= element
}

func (i Interval) splitOn(limit int) (Interval, Interval) {
	return Interval{i.lower, limit}, Interval{limit, i.upper}
}

func (i Interval) differenceFrom(other Interval) []Interval {
	if Intersect(i, other).IsEmpty() {
		return []Interval{i}
	}
	if other.contains(i) {
		return []Interval{}
	}
	if i.contains(other) {
		return []Interval{{i.lower, other.lower}, {other.upper, i.upper}}
	}
	if i.containsElement(other.lower) {
		return []Interval{{i.lower, other.lower}}
	}
	if i.containsElement(other.upper) {
		return []Interval{{other.upper, i.upper}}
	}
	panic("The impossible has occurred")
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
