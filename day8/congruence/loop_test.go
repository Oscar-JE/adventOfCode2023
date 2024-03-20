package congruence

import (
	"testing"
)

func TestCutWithUnitLoop(t *testing.T) {
	loop := InitTailedLoop(3, 4, []int{0, 1, 2, 3})
	loop2 := loop.Synk(loop)
	if !(loop.Eq(loop2)) {
		t.Errorf("Synk with self not returning self")
	}
}
