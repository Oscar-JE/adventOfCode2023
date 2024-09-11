package environment

import "math"

type StateCache struct { //testar debba om det krisar
	statePaths []int
	rows       int
}

func InitStateCach(rows int, cols int) StateCache {
	maxHash := rows * cols * len(directions) * maxStepInline
	cachedValues := []int{}
	for range maxHash {
		cachedValues = append(cachedValues, math.MaxInt)
	}
	return StateCache{statePaths: cachedValues, rows: rows}
}

func (c *StateCache) Set(s State, pathV int) {
	index := s.hashCode(c.rows)
	c.statePaths[index] = pathV
}

func (c StateCache) Get(s State) int {
	return c.statePaths[s.hashCode(c.rows)]
}
