package solution

import "math"

type Race struct {
	Time     int
	Distance int
}

func boarders(r Race) Segment {
	T := float64(r.Time)
	D := float64(r.Distance)
	lower := T/2 - math.Sqrt(math.Pow(T/2, 2)-D)
	upper := T/2 + math.Sqrt(math.Pow(T/2, 2)-D)
	return Segment{lower: lower, upper: upper}
}

func NrOfWins(r Race) int {
	seg := boarders(r)
	return seg.NrOfIntsInside()
}
