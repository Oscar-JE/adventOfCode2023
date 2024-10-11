package hole

import (
	"day16/vec"
	"day18/segment"
	"fmt"
	"testing"
)

func TestFindTopSegment(t *testing.T) {
	segs := []segment.Segment{segment.Init(vec.Init(-1, -1), 2, vec.Init(0, 1)),
		segment.Init(vec.Init(-1, 1), 2, vec.Init(1, 0)),
		segment.Init(vec.Init(1, 1), 2, vec.Init(0, -1)),
		segment.Init(vec.Init(1, -1), 2, vec.Init(-1, 0))}
	h := Init(segs)
	topSeg := h.findTopSegment()
	expected := segment.Init(vec.Init(-1, -1), 2, vec.Init(0, 1))
	if topSeg != expected {
		t.Errorf("examine top segment")
	}
}

func TestFindInterior(t *testing.T) {
	segments := []segment.Segment{segment.Init(vec.Init(0, 0), 2, vec.Init(0, 1)), segment.Init(vec.Init(0, 2), 2, vec.Init(1, 0)), segment.Init(vec.Init(2, 2), 2, vec.Init(0, -1)), segment.Init(vec.Init(2, 0), 2, vec.Init(-1, 0))}
	hole := Init(segments)
	interiorStartingPoints := hole.findStartingPoints()
	if len(interiorStartingPoints) != 1 {
		t.Errorf("Lenght of interior poins for box case should be one but was %d", len(interiorStartingPoints))
	}
	interiorPoint := interiorStartingPoints[0]
	if interiorPoint != vec.Init(1, 1) {
		t.Errorf("wrong internal starting point")
	}
}

func TestLengthOfBoarder(t *testing.T) {
	segments := []segment.Segment{
		segment.Init(vec.Init(0, 0), 2, vec.Init(0, 1)),
		segment.Init(vec.Init(0, 2), 2, vec.Init(1, 0)),
		segment.Init(vec.Init(2, 2), 2, vec.Init(0, -1)),
		segment.Init(vec.Init(2, 0), 2, vec.Init(-1, 0))}
	h := Init(segments)
	res := h.lengthOfBoarder()
	expected := 8
	if res != expected {
		t.Errorf("Expeted boarder length of %d in square with one interior point", expected)
	}
}

func TestLengthOfBoarderShortExample(t *testing.T) {
	h := InitShortExample()
	len := h.lengthOfBoarder()
	expected := 38
	if len != expected {
		t.Errorf("lengt should be %d but was %d", expected, len)
	}
}

func TestNrInteriorPoints(t *testing.T) {
	h := InitShortExample()
	interior := h.NrOfInteriorPoints2()
	expected := 24
	if interior != expected {
		t.Errorf("We expect %d interior points bur we got %d", expected, interior)
	}
}

func TestExcavatedPoint(t *testing.T) {
	h := InitShortExample()
	res := h.NrExcavated()
	expected := 62
	if res != expected {
		t.Errorf("Nr of excavated whould be %d but was %d for short example", expected, res)
	}
}

func TestInitShortExample(t *testing.T) {
	t.Skip()
	h := InitShortExample()
	fmt.Println(h)
}

func InitShortExample() Hole { // nu får vi ett off by one error
	segments := []segment.Segment{ // antagligen är bara den här uppsättningen fel
		segment.Init(vec.Init(0, 0), 6, vec.Init(0, 1)),
		segment.Init(vec.Init(0, 6), 5, vec.Init(1, 0)),
		segment.Init(vec.Init(5, 6), 2, vec.Init(0, -1)),
		segment.Init(vec.Init(5, 4), 2, vec.Init(1, 0)),
		segment.Init(vec.Init(7, 4), 2, vec.Init(0, 1)),
		segment.Init(vec.Init(7, 6), 2, vec.Init(1, 0)),
		segment.Init(vec.Init(9, 6), 5, vec.Init(0, -1)),
		segment.Init(vec.Init(9, 1), 2, vec.Init(-1, 0)),
		segment.Init(vec.Init(7, 1), 1, vec.Init(0, -1)),
		segment.Init(vec.Init(7, 0), 2, vec.Init(-1, 0)),
		segment.Init(vec.Init(5, 0), 2, vec.Init(0, 1)),
		segment.Init(vec.Init(5, 2), 3, vec.Init(-1, 0)),
		segment.Init(vec.Init(2, 2), 2, vec.Init(0, -1)),
		segment.Init(vec.Init(2, 0), 2, vec.Init(-1, 0))}
	return Init(segments)
}
