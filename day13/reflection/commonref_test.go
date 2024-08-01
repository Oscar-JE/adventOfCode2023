package reflection

import (
	"testing"
)

func TestIsReflectionindex(t *testing.T) {
	values := []int{0, 0}
	if !isReflektionIndex(0, values) {
		t.Errorf("rethink isReflektion")
	}
	if isReflektionIndex(1, values) {
		t.Errorf("Can never have an mirror outside the field")
	}
}

func TestIsReflectionindexLongExample(t *testing.T) {
	values := []int{1, 1, 0, 1, 1, 0, 1}
	if !isReflektionIndex(3, values) {
		t.Errorf("Examin the long case for isReflectionIndex")
	}
	if !isReflektionIndex(0, values) {
		t.Errorf("Examin the long case for isReflectionIndex")
	}
}

func TestConstruct(t *testing.T) {
	values := []int{0, 0}
	refConsummer := FindPossibleRef(values)
	if refConsummer.refIndex[0] != 0 {
		t.Errorf("wrong construction")
	}
}

func TestAlign(t *testing.T) {
	havRef := []int{1, 1}
	noRef := []int{0, 1}
	refC := FindPossibleRef(havRef)
	refC.Align(havRef)
	if refC.refIndex[0] != 0 {
		t.Errorf("Align not matching")
	}
	refC.Align(noRef)
	if len(refC.refIndex) != 0 {
		t.Errorf("on no match should set state to -1")
	}
}
