package set

// Skulle kunna implementera som ett bin tree men känner att
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

func (s *Set) add(seg Interval) {

	newDisjunctSegments := []Interval{}
	segIsAdded := false

	for _, oldSeg := range s.disjunctIntervals {
		if oldSeg.isLowerThan(seg) {
			newDisjunctSegments = append(newDisjunctSegments, oldSeg)
			continue
		} else if !Intersect(oldSeg, seg).IsEmpty() {
			seg = seg.unionOfConnected(oldSeg)
			continue
		} else if seg.isLowerThan(oldSeg) {
			if !segIsAdded {
				newDisjunctSegments = append(newDisjunctSegments, seg)
				segIsAdded = true
			}
			newDisjunctSegments = append(newDisjunctSegments, oldSeg)
			continue
		}
	}

	if !segIsAdded {
		newDisjunctSegments = append(newDisjunctSegments, seg)
		segIsAdded = true
	}

	s.disjunctIntervals = newDisjunctSegments
}

func (s Set) Eq(other Set) bool {
	if len(s.disjunctIntervals) != len(other.disjunctIntervals) {
		return false
	}
	if len(s.disjunctIntervals) == 0 {
		return true
	}
	for i := range s.disjunctIntervals {
		if s.disjunctIntervals[i] != other.disjunctIntervals[i] {
			return false
		}
	}
	return true
}

func Union(s1 Set, s2 Set) Set {
	for _, interval := range s1.disjunctIntervals {
		s2.add(interval)
	}
	return s2
}
