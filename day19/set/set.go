package set

type Set struct {
	disjunctIntervals []Interval
}

func Init(start int, end int) Set {
	return Set{[]Interval{{start, end}}}
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

func (s Set) DemandAbove(limit int) (Set, Set) {
	newDisjunctSegments := []Interval{}
	newDisjunctSegmentComplement := []Interval{}
	for _, seg := range s.disjunctIntervals {
		if seg.isLowerThanEl(limit) {
			newDisjunctSegments = append(newDisjunctSegments, seg)
			continue
		}
		if seg.isLargerThanEl(limit) {
			newDisjunctSegmentComplement = append(newDisjunctSegmentComplement, seg)
		}
		if seg.containsElement(limit) {
			lowerInterval, upperInterval := seg.splitOn(limit)
			newDisjunctSegments = append(newDisjunctSegments, lowerInterval)
			newDisjunctSegmentComplement = append(newDisjunctSegmentComplement, upperInterval)
		}
	}
	return Set{newDisjunctSegments}, Set{newDisjunctSegmentComplement}
}

func (s Set) DemandBelow(limit int) (Set, Set) {
	newDisjunctSegments := []Interval{}
	newDisjunctSegmentComplement := []Interval{}
	for _, seg := range s.disjunctIntervals {
		if seg.isLowerThanEl(limit) {
			newDisjunctSegmentComplement = append(newDisjunctSegmentComplement, seg)
		}
		if seg.isLargerThanEl(limit) {
			newDisjunctSegments = append(newDisjunctSegments, seg)
		}
		if seg.containsElement(limit) {
			lowerInterval, upperInterval := seg.splitOn(limit)
			newDisjunctSegments = append(newDisjunctSegments, upperInterval)
			newDisjunctSegmentComplement = append(newDisjunctSegmentComplement, lowerInterval)
		}
		newDisjunctSegments = append(newDisjunctSegments, seg.DemandBelow(limit))
	}
	return Set{newDisjunctSegments}, Set{newDisjunctSegmentComplement}
}
