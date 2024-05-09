package partone

import (
	"day12/field"
	"day12/order"
	"testing"
)

func TestCaseInstruktion1(t *testing.T) {
	f := field.Init("???.###")
	o := order.Init([]int{1, 1, 3})
	res := NumberOfCombination(f, o)
	if res != 1 {
		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara 1", res)
	}
}

func TestCaseInstruktionGotcha(t *testing.T) {
	f := field.Init("???.###.#")
	o := order.Init([]int{1, 1, 3})
	res := NumberOfCombination(f, o)
	if res != 0 {
		t.Errorf("Gotcha var %d . Men borde vara 0", res)
	}
}

func TestCaseInstruktion2(t *testing.T) {
	f := field.Init(".??..??...?##.")
	o := order.Init([]int{1, 1, 3})
	res := NumberOfCombination(f, o)
	if res != 4 {
		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara %d", res, 4)
	}
}

func TestCaseInstruktion3(t *testing.T) {
	f := field.Init("?#?#?#?#?#?#?#?")
	o := order.Init([]int{1, 3, 1, 6})
	res := NumberOfCombination(f, o)
	expected := 1
	if res != expected {
		t.Errorf("första testcaset från instruktionen failede res var %d . Men borde vara %d", res, expected)
	}
}
