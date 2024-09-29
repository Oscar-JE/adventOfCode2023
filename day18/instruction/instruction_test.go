package instruction

import (
	"day16/vec"
	"testing"
)

func TestParse(t *testing.T) {
	rep := "R 6 (#70c710)"
	expected := Instruction{direction: vec.Init(0, 1), distance: 6}
	res := parse(rep)
	if res != expected {
		t.Errorf(rep + " is not parsed correctly")
	}
}
