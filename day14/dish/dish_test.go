package dish

import (
	"day14/dish/pice"
	"testing"
)

func TestInitFromText(t *testing.T) {
	rep := ".O\n#.\n.."
	d := InitFromText(rep)
	if d.layout.Get(0, 0) != pice.Init(pice.Empty) || d.layout.Get(0, 1) != pice.Init(pice.Movable) || d.layout.Get(1, 0) != pice.Init(pice.Fixed) || d.layout.Get(1, 1) != pice.Init(pice.Empty) || d.layout.Get(2, 0) != pice.Init(pice.Empty) || d.layout.Get(2, 1) != pice.Init(pice.Empty) {
		t.Errorf("InitFrom text not working correctly")
	}
}
