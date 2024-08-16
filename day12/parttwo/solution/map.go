package solution

type Counter struct {
	m map[string]int
}

func (c Counter) Add(s SolutionKey, solutions int) {
	c.m[s.String()] = solutions
}

func (c Counter) Get(s SolutionKey) int {
	val, ok := c.m[s.String()]
	if ok {
		return val
	}
	return 0
}
