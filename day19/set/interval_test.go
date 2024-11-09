package set

import "testing"

func TestDifferenceNoIntersection(t *testing.T) {
	seg1 := Interval{0, 2}
	seg2 := Interval{2, 4}
	res := seg1.differenceFrom(seg2)
	if len(res) != 1 {
		t.Errorf("diff ska ha längd 1 i ingen intersect caset")
	}
	resEl := res[0]
	if resEl != seg1 {
		t.Errorf("ingen intersection ska leda till inget avdrag")
	}
}

func TestDiffferenceIntersectionAbove(t *testing.T) {
	seg1 := Interval{0, 2}
	seg2 := Interval{1, 4}
	res := seg1.differenceFrom(seg2)
	if len(res) != 1 {
		t.Errorf("With a intersection not entire segment should have one element")
	}
	resEl := res[0]
	expected := Interval{0, 1}
	if resEl != expected {
		t.Errorf("Should result in segment (0,1)")
	}
}

func TestDifferenceIntersectionBellow(t *testing.T) {
	seg1 := Interval{0, 2}
	seg2 := Interval{-1, 1}
	res := seg1.differenceFrom(seg2)
	if len(res) != 1 {
		t.Errorf("blä")
	}
	resEl := res[0]
	expected := Interval{1, 2}
	if resEl != expected {
		t.Errorf("assertion error")
	}
}

func TestSecondContains(t *testing.T) {
	seg1 := Interval{0, 1}
	seg2 := Interval{0, 1}
	res := seg1.differenceFrom(seg2)
	if len(res) != 0 {
		t.Errorf("meditate on a set diffed with itself")
	}
}

func TestFirstContainsInner(t *testing.T) {
	seg1 := Interval{0, 10}
	seg2 := Interval{1, 5}
	res := seg1.differenceFrom(seg2)
	if len(res) != 2 {
		t.Errorf(" If interior point is removed there is points to the left and right")
		return
	}
	resEl1 := res[0]
	resEl2 := res[1]
	expected1 := Interval{0, 1}
	expected2 := Interval{5, 10}
	if resEl1 != expected1 && resEl2 != expected2 {
		t.Errorf("meditate on the removal of interior segment from larger segment")
	}
}
