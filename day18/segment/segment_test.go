package segment

import (
	"day16/vec"
	"day18/instruction"
	"testing"
)

func TestInOutside(t *testing.T) {
	seg := InitFromInstruction(vec.Init(2, 0), instruction.Init(2, vec.Init(0, 1)))
	p := vec.Init(5, 5)
	if seg.In(p) {
		t.Errorf("(5,5) is not in given segment")
	}
}

func TestInStart(t *testing.T) {
	seg := InitFromInstruction(vec.Init(2, 0), instruction.Init(2, vec.Init(0, 1)))
	p := vec.Init(2, 0)
	if seg.In(p) {
		t.Errorf("Start of segment should not be included in segment")
	}
}

func TestInEnd(t *testing.T) {
	seg := InitFromInstruction(vec.Init(3, 3), instruction.Init(2, vec.Init(-1, 0)))
	end := vec.Init(1, 3)
	if !seg.In(end) {
		t.Errorf("End should be included in segment")
	}
}

func TestInMiddle(t *testing.T) {
	seg := InitFromInstruction(vec.Init(3, 3), instruction.Init(2, vec.Init(-1, 0)))
	p := vec.Init(2, 3)
	if !seg.In(p) {
		t.Errorf("Investigate case where point is in the middle of segment")
	}
}
