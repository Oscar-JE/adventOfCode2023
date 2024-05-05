package field

import "testing"

func TestTrivial(t *testing.T) {
	f := Init("#")
	res := len(f.PossiblePlacements(1))
	if res != 1 {
		t.Errorf("Något är brutalt fel")
	}
}

func TestSmpleBaseCase(t *testing.T) {
	f := Init("?.?")
	res := len(f.PossiblePlacements(1))
	if res != 2 {
		t.Errorf("nummber of combination returned wrong in the case ?.? 1")
	}
}

func TestRestriction(t *testing.T) {
	f := Init("#.?")
	res := len(f.PossiblePlacements(1))
	if res != 1 {
		t.Errorf("dags att läsa instructionerna igen #.? 1 , ska vara 1 var  med var %d", res)
	}
}
