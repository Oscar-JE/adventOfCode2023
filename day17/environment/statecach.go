package environment

import "math"

type StateCache struct { //testar debba om det krisar
	statePaths []int
	rows       int
}

func InitStateCach(state State, rows int, cols int) StateCache {
	maxHash := state.MaxHash(rows, cols)
	cachedValues := []int{}
	for range maxHash {
		cachedValues = append(cachedValues, math.MaxInt)
	}
	return StateCache{statePaths: cachedValues, rows: rows}
}

func (c *StateCache) SetValue(s State, pathV int) {
	index := s.hashCode(c.rows)
	c.statePaths[index] = pathV
}

func (c StateCache) GetValue(s State) int {
	return c.statePaths[s.hashCode(c.rows)]
}

func (c StateCache) Has(s State) bool {
	return c.GetValue(s) < math.MaxInt
}

func (c StateCache) ProcentageFilled() float64 {
	numberOfChached := 0
	for _, cost := range c.statePaths {
		if cost < math.MaxInt {
			numberOfChached++
		}
	}
	return float64(numberOfChached) / float64(len(c.statePaths))
}
