package game

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (s Set) contains(other Set) bool {
	return other.Red <= s.Red && other.Green <= s.Green && other.Blue <= s.Blue
}

func (s *Set) union(other Set) {
	s.Red = max(s.Red, other.Red)
	s.Green = max(s.Green, other.Green)
	s.Blue = max(s.Blue, other.Blue)
}

func (s Set) Power() int {
	return s.Red * s.Green * s.Blue
}

func FindMinimum(sets []Set) Set {
	set1 := sets[0]
	for i := 1; i < len(sets); i++ {
		set1.union(sets[i])
	}
	return set1
}
