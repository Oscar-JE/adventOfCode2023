package box

import (
	"testing"
)

func TestAddNoDupplicate(t *testing.T) {
	box := Init([]int{1, 2, 3})
	box.Add(4)
	if box.lenses[3] != 4 {
		t.Errorf("Add failed with no dupplicate")
	}
}

func TestAddWithDuplicate(t *testing.T) {
	box := Init([]int{1, 2, 3})
	box.Add(2)
	if box.lenses[0] != 1 || box.lenses[1] != 2 || box.lenses[2] != 3 {
		t.Errorf("Add failed with duplicate")
	}
}

func TestRemoveNoMatch(t *testing.T) {
	box := Init([]int{1, 2})
	box.Remove(10)
	if box.lenses[0] != 1 || box.lenses[1] != 2 {
		t.Errorf("remove removed even without match")
	}
}

func TestRemoveWithMultipleMatch(t *testing.T) {
	box := Init([]int{1, 2, 1})
	box.Remove(1)
	if box.lenses[0] != 2 || box.lenses[1] != 1 {
		t.Errorf("examine remove when there is multiple matches")
	}
}
