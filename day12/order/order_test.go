package order

import "testing"

func TestRolingRsizeTrivial(t *testing.T) {
	o := Init([]int{})
	res := o.RolingRsize(0, 0)
	if !o.Eq(res) {
		t.Errorf("RolingRsize(0,0) is not a Identity mapping")
	}
}

func TestRolingRsizeExpand(t *testing.T) {
	o := Init([]int{1, 2, 3})
	res := o.RolingRsize(-1, 1)
	expected := Init([]int{3, 1, 2, 3, 1})
	if !res.Eq(expected) {
		t.Errorf("RolingRsize(-1,1)")
	}
}

func TestRolingRsizeContract(t *testing.T) {
	o := Init([]int{1, 2, 3})
	res := o.RolingRsize(1, -1)
	expected := Init([]int{2})
	if !res.Eq(expected) {
		t.Errorf("RolingRsize(1,-11)")
	}
}

func TestRolingRsizeIdentity(t *testing.T) {
	o := Init([]int{1, 2, 3})
	res1 := o.RolingRsize(6, 6)
	res2 := o.RolingRsize(-9, -9)
	if !res1.Eq(res2) && res2.Eq(o) {
		t.Errorf("RolingRsizeIdentity")
	}
}

func TestSplitIndex(t *testing.T) {
	o := Init([]int{1, 2, 3})
	lowO, uppO := o.SubOrder(0)
	lowExp, uppExp := Init([]int{1}), Init([]int{2, 3})
	if lowExp.Eq(lowO) && uppExp.Eq(uppO) {
		t.Errorf("fan")
	}
}
