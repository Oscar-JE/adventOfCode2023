package ascii

import "testing"

func TestAsciiH(t *testing.T) {
	ascii_tester('H', 72, t)
	ascii_tester('A', 65, t)
	ascii_tester('S', 83, t)
}

func ascii_tester(r rune, v int, t *testing.T) {
	if Value(r) != v {
		t.Errorf(" %c should be %d", r, v)
	}
}
