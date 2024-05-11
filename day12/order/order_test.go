package order

import "testing"

func TestRollBeginningHappy(t *testing.T) {
	order := Init([]int{1, 2, 3})
	expected := Init([]int{3})
	res := order.RollBeginning(2)
	if !res.Eq(expected) {
		t.Errorf("review roll beginning for order")
	}
}

func TestRollBegginingNegativ(t *testing.T) {
	order := Init([]int{1, 2, 3})
	expecteed := Init([]int{3, 1, 2, 3, 1, 2, 3})
	res := order.RollBeginning(-4)
	if !res.Eq(expecteed) {
		t.Errorf("roll left beter sig konstigt")
	}
}

func TestRollEnd(t *testing.T) {
	order := Init([]int{1, 2, 3})
	expected := Init([]int{1, 2, 3, 1, 2, 3, 1})
	res := order.RollEnd(4)
	if !res.Eq(expected) {
		t.Errorf("review roll end for order")
	}
}

func TestRollEndNegative(t *testing.T) {
	order := Init([]int{1, 2, 3})
	expected := Init([]int{1})
	res := order.RollEnd(-2)
	if !res.Eq(expected) {
		t.Errorf("review expand right for order")
	}
}
