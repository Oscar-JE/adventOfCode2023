package set

//Skulle kunna implementera som ett bin tree men känner att 
// det skulle vara mer arbeta än nödvändigt såhär innan enkla exemplet gått igenom
type Set struct {
	disjunctIntervals []Interval
}

func (s Set) Cardinality() int {
	sum := 0
	for _, seg := range s.disjunctIntervals {
		sum += seg.Cardinality()
	}
	return sum
}

func (s Set) differenceFrom(seg Interval) []Interval {
	return []Interval{seg}
}
