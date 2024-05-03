package springs

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	o := Operational()
	b := Damaged()
	u := Unknown()
	res := fmt.Sprint(o, b, u)
	if res != ". # ?" {
		t.Errorf("examin Conditions String method")
	}
}
