package vec

import "testing"

func TestRotate(t *testing.T) {
	v := Init(2, 1)
	r1 := v.Rotate90DegreesUp()
	r2 := r1.Rotate90DegreesUp()
	r3 := r2.Rotate90DegreesUp()
	r4 := r3.Rotate90DegreesUp()
	// identity
	if r4 != v {
		t.Errorf("Four rotations did not form an identity")
	}
	if r1 != Init(-1, 2) {
		t.Errorf("check one rotation")
	}
	if r2 != Init(-2, -1) {
		t.Errorf("check two rotations")
	}
	if r3 != Init(1, -2) {
		t.Errorf("check three rotations")
	}
}
