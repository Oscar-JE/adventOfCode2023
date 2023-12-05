package game

import (
	"testing"
)

func TestContainsTrivial(t *testing.T) {
	set1 := Set{0, 0, 0}
	set2 := Set{1, 1, 1}
	if set1.contains(set2) {
		t.Errorf("TrivialTest: set1 contains set2. That should not be the case ")
	}
	if !set2.contains(set1) {
		t.Errorf("TrivialTest: Set2 should contain set1")
	}
}
