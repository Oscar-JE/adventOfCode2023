package point

import "testing"

func TestFindClosestToPoint(t *testing.T) {
	vList := InitVecList([]Vec{Init(3, 1), Init(3, 2), Init(3, 3)})
	vec, index := vList.FindCLosestTo(Init(0, 0))
	if vec != vList.vecs[0] {
		t.Errorf("closest Vektor not correct")
	}
	if index != 0 {
		t.Errorf("argmin in find closes has not worked")
	}
}

func TestFindClosestToPointDebuggExemple(t *testing.T) {
	vList := InitVecList([]Vec{Init(1, 1), Init(2, 1), Init(3, 1), Init(3, 2), Init(3, 3), Init(2, 3), Init(1, 3), Init(1, 2)})
	vec, index := vList.FindCLosestTo(Init(0, 0))
	if vec != vList.vecs[0] {
		t.Errorf("closest Vektor not correct")
	}
	if index != 0 {
		t.Errorf("argmin in find closes has not worked")
	}
}
