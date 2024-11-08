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
	//todo fortsätt här
}
