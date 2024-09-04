package environment

import (
	"day16/vec"
	"testing"
)

func TestNextStatesStepsNotLimiting(t *testing.T) {
	start := Init(vec.Init(0, 0), north, 0)
	nextStates := start.NextPossibleStates()
	if len(nextStates) != 3 {
		t.Errorf("length of next steps should be 3 ig number of steps take in less then 3")
	}
	s0 := Init(vec.Init(-1, 0), north, 1)
	s1 := Init(vec.Init(0, 1), east, 0)
	s2 := Init(vec.Init(0, -1), west, 0)
	if nextStates[0] != s0 || nextStates[1] != s1 || nextStates[2] != s2 {
		t.Errorf("slice of next states not what expected in next states without step limit")
	}
}

func TestNextStatesStepsLimitingFactor(t *testing.T) {
	start := Init(vec.Init(0, 0), north, 2) // detta bör vara rätt
	nextStates := start.NextPossibleStates()
	if len(nextStates) != 2 {
		t.Errorf("length of next steps should be 2 ")
	}
	s0 := Init(vec.Init(0, 1), east, 0)
	s1 := Init(vec.Init(0, -1), west, 0)
	if nextStates[0] != s0 || nextStates[1] != s1 {
		t.Errorf("slice of next states not what expected in next states with step limit")
	}
}
