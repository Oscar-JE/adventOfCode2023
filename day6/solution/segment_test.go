package solution

import (
	"testing"
)

func TestNrOfIntsInside(t *testing.T) {
	seg := Segment{lower: 1.5, upper: 2.0}
	res := seg.NrOfIntsInside()
	if res != 0 {
		t.Errorf("There is no integer inside the segment %v ", seg)
	}
}

func TestNrOfIntsInside2(t *testing.T) {
	seg := Segment{lower: 1.0, upper: 5.0}
	res := seg.NrOfIntsInside()
	if res != 3 {
		t.Errorf("There is 3 integer inside the segment %v ", seg)
	}
}
