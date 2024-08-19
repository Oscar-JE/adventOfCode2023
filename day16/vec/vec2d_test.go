package vec

import "testing"

func TestTurningUppZero(t *testing.T) { // kommer jag behöva funktionen?
	v := Init(0, 0)
	r := Trun90upp(v)
	if v != r {
		t.Errorf("turning of zero should be zero´")
	}
}
