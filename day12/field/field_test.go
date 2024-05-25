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

func TestTakeLeftTrival(t *testing.T) {
	f := Init("?..")
	res, ok := f.RestritLeft(-1)
	if !ok {
		t.Errorf("wrong restriktion control")
	}
	expected := Init("?..")
	if !res.Eq(expected) {
		t.Errorf("wrong expected output")
	}
}

func TestTakeLeftZero(t *testing.T) {
	f := Init("?..")
	res, ok := f.RestritLeft(0)
	if !ok {
		t.Errorf("wrong restriktion control")
	}
	expected := Init("...")
	if !res.Eq(expected) {
		t.Errorf("wrong expected output")
	}
}

func TestTakeLeftOne(t *testing.T) {
	f := Init("?..")
	res, ok := f.RestritLeft(1)
	expected := Init("#..")
	if !res.Eq(expected) || !ok {
		t.Errorf("investigate TakeLefOne")
	}
}

func TestTakeLeftUnHappy(t *testing.T) {
	noSpace := Init("###")
	impossibleBroken := Init(".#.")
	_, takeFromNoSpace := noSpace.RestritLeft(0)
	if takeFromNoSpace {
		t.Errorf("impossible take")
	}
	_, stumble := impossibleBroken.RestritLeft(2)

	if stumble {
		t.Errorf("impossible take left")
	}
}
