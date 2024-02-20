package direction

import (
	"testing"
)

func TestDirectionName(t *testing.T) {
	dir := direction.Left
	res := dit.GetName()
	if res != "L"{
		t.Errorf("res was ")
	}
}
