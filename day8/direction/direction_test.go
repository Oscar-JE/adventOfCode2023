package direction

import (
	"testing"
)

func TestDirectionName(t *testing.T) {
	dir := LEFT
	res := dir.String()
	if res != "L" {
		t.Errorf("res was ")
	}
}
