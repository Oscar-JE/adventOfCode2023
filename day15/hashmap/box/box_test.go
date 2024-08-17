package box

import (
	"day15/hashmap/lens"
	"testing"
)

func TestAddNoDupplicate(t *testing.T) {
	box := Init([]lens.Lens{lens.Init("1", 1), lens.Init("2", 2), lens.Init("3", 3)})
	box.Add(lens.Init("4", 4))
	if box.lenses[3] != lens.Init("4", 4) {
		t.Errorf("Add failed with no dupplicate")
	}
}

func TestAddWithDuplicate(t *testing.T) {
	box := Init([]lens.Lens{lens.Init("1", 1), lens.Init("2", 2), lens.Init("3", 3)})
	box.Add(lens.Init("2", 2))
	if box.lenses[0] != lens.Init("1", 1) || box.lenses[1] != lens.Init("2", 2) || box.lenses[2] != lens.Init("3", 3) {
		t.Errorf("Add failed with duplicate")
	}
}

func TestRemoveNoMatch(t *testing.T) {
	box := Init([]lens.Lens{lens.Init("1", 1), lens.Init("2", 2)})
	box.Remove(lens.Init("10", 10))
	if box.lenses[0] != lens.Init("1", 1) || box.lenses[1] != lens.Init("2", 2) {
		t.Errorf("remove removed even without match")
	}
}

func TestRemoveWithMultipleMatch(t *testing.T) {
	box := Init([]lens.Lens{lens.Init("1", 1), lens.Init("2", 2), lens.Init("1", 1)})
	box.Remove(lens.Init("1", 1))
	if box.lenses[0] != lens.Init("2", 2) || box.lenses[1] != lens.Init("1", 1) {
		t.Errorf("examine remove when there is multiple matches")
	}
}

func TestFocusingPower(t *testing.T) {
	box := Init([]lens.Lens{lens.Init("rn", 1), lens.Init("cm", 2)})
	res := box.FocusingPower()
	expected := 5
	if res != expected {
		t.Errorf("examine focusing power for Box 0")
	}
}
