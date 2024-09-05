package environment

import (
	"day16/vec"
	"testing"
)

func TestNexStates(t *testing.T) {
	env := InitEnv([]int{1, 2, 3, 4}, 2, 2)
	state := Init(vec.Init(0, 0), vec.Init(-1, 0), 0)
	nextStates := env.NextStates(state)
	expected := Init(vec.Init(0, 1), vec.Init(0, 1), 0)
	if len(nextStates) != 1 || nextStates[0] != expected {
		t.Errorf("unexpected state in TestNextStates")
	}
}
