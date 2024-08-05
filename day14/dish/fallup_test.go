package dish

import (
	"day14/dish/pice"
	"testing"
)

func TestFalluppOne(t *testing.T) {
	pices := []pice.Pice{pice.Empty, pice.Movable, pice.Fixed, pice.Empty, pice.Movable}
	lowest := fallUppOne(1, pices)
	highest := fallUppOne(4, pices)
	if lowest != 0 || highest != 3 {
		t.Errorf("FallUpp one behaves unexpectedly")
	}
}

func TestFallUpp(t *testing.T) {
	pices := []pice.Pice{pice.Empty, pice.Movable, pice.Fixed, pice.Empty, pice.Movable}
	expected := []pice.Pice{pice.Movable, pice.Empty, pice.Fixed, pice.Movable, pice.Empty}
	res := fallUpp(pices)
	for i, el := range expected {
		if el != res[i] {
			t.Errorf("examin fallUpp")
		}
	}
}
