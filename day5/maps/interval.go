package maps

type Interval struct {
	from   int
	length int
}

func (i Interval) Contains(element int) bool {
	return i.from <= element && element < i.from+i.length
}

func InitInterval(from int, length int) Interval {
	return Interval{from: from, length: length}
}

type Set struct {
	subSets []Interval
}

func InitSeedSet(intervals []Interval) Set {
	return Set{subSets: intervals}
}

func (s Set) Contains(element int) bool {
	c := false
	for _, subSet := range s.subSets {
		c = c || subSet.Contains(element)
	}
	return c
}
