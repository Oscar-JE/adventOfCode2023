package policy

import (
	"day8/direction"
	"testing"
)

func TestPolicy(t *testing.T) {
	pol := Init("RL")
	d1 := pol.GetNext()
	d2 := pol.GetNext()
	d3 := pol.GetNext()
	if d1 != direction.RIGHT {
		t.Errorf("wrong first direction")
	}
	if d2 != direction.LEFT {
		t.Errorf("wrong second direction")
	}
	if d3 != direction.RIGHT {
		t.Errorf("wrong third direction")
	}
}
