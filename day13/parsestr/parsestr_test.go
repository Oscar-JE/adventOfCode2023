package parsestr

import (
	"testing"
)

func TestStringParse(t *testing.T) {
	str := "#.##\n..#."
	m := ParseFromString(str)
	if m.Get(1, 0) != 0 || m.Get(0, 0) != 1 {
		t.Errorf("debugg String parse")
	}
}
