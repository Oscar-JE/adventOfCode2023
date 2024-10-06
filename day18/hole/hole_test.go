package hole

import (
	"day16/vec"
	"testing"
)

func TestFindTopNeighbors(t *testing.T) {
	h := Init([]vec.Vec2d{vec.Init(-1, 0), vec.Init(-1, 1), vec.Init(0, 1), vec.Init(0, 0)})
	v1, v2 := h.findIndexOfPairadjacentToTopBoarder()
	if v1 != vec.Init(-1, 0) || v2 != vec.Init(-1, 1) {
		t.Errorf("wrong neighbors found")
	}
}

func TestFindInterior(t *testing.T) {
	points := []vec.Vec2d{vec.Init(0, 0), vec.Init(0, 1), vec.Init(0, 2), vec.Init(1, 2), vec.Init(2, 2), vec.Init(2, 1), vec.Init(2, 0), vec.Init(1, 0)}
	hole := Init(points)
	interiorStartingPoints := hole.findStartingPoints()
	if len(interiorStartingPoints) != 1 {
		t.Errorf("Lenght of interior poins for box case should be one but was %d", len(interiorStartingPoints))
	}
	interiorPoint := interiorStartingPoints[0]
	if interiorPoint != vec.Init(1, 1) {
		t.Errorf("wrong internal starting point")
	}
}
