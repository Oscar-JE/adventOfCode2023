package dijkstras

import "testing"

// test implementations

type testState struct {
	rep int
}

func (s testState) GetRep() int {
	return s.rep
}

type testEnvironment struct {
	transferMatrix [][]int
}

func (e testEnvironment) TransFer(t testState) []StateAndCost[testState] {
	row := e.transferMatrix[t.GetRep()]
	statesAndCosts := []StateAndCost[testState]{}
	for index, cost := range row {
		if cost < 0 {
			continue
		}
		statesAndCosts = append(statesAndCosts, StateAndCost[testState]{testState{index}, cost})
	}
	return statesAndCosts
}

type TestCach[E comparable] struct {
	cach []StateAndCost[E]
}

func (t *TestCach[E]) Has(s E) bool {
	for _, el := range t.cach {
		if el.state == s {
			return true
		}
	}
	return false
}

func (t *TestCach[E]) SetValue(s E, val int) {
	for index, el := range t.cach {
		if el.state == s {
			t.cach[index].cost = val
			return
		}
	}
	t.cach = append(t.cach, StateAndCost[E]{s, val})
}

func (t *TestCach[E]) GetValue(s E) int {
	for _, el := range t.cach {
		if el.state == s {
			return el.cost
		}
	}
	return 100
}

func (t TestCach[E]) ProcentageFilled() float64 {
	return 0
}

func TestDijkstras(t *testing.T) {
	env := testEnvironment{[][]int{
		[]int{-1, 5, 1, -1},
		[]int{-1, -1, -1, 1},
		[]int{-1, 2, -1, 7},
		[]int{-1, -1, -1, -1}}}
	startState := testState{0}
	cach := &TestCach[testState]{[]StateAndCost[testState]{}}
	dij := Init[testState](env, cach)
	dists := dij.SmallestDist([]testState{testState{3}}, startState, 0)
	if dists != 4 {
		t.Errorf("distance for test environment should be 4, but was %d", dists)
	}
}
