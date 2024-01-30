package congruence

import (
	"testing"
)

func TestCutWithUnitLoop(t *testing.T) {
	t.Skip("working on this later")
	unit := TailedLoop{tail: 0, loopLength: 1, winningPoint: 0}
	anyLoop := TailedLoop{tail: 23, loopLength: 450, winningPoint: 55}
	cutLoop, _ := cut(anyLoop, unit)
	if anyLoop != cutLoop {
		t.Errorf("cut with unit loop should not change the loop")
	}
}

func TestEmptyCutNoTail(t *testing.T) {
	t.Skip("working on this later")
	l1 := TailedLoop{tail: 0, loopLength: 2, winningPoint: 0}
	l2 := TailedLoop{tail: 0, loopLength: 2, winningPoint: 1}
	_, err := cut(l1, l2)
	if err == nil {
		t.Errorf("An cut that would return the empty loop should return an error")
	}
}
