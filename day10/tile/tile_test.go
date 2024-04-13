package tile

import (
	"day10/point"
	"testing"
)

func TestInitEmpty(t *testing.T) {
	empty := Init('.', point.Init(0, 0))
	set := empty.Ways()
	if !set.IsEmpty() {
		t.Errorf("empty not behaving correctly")
	}
}

func TestInitStraigthStanding(t *testing.T) {
	straigth := Init('|', point.Init(0, 1))
	res := straigth.Ways()
	expected := point.InitVecSet([]point.Vec{point.Init(-1, 1), point.Init(1, 1)})
	if !expected.Eq(res) {
		t.Errorf("kolla närmare på ways av stående streck")
	}
}

func TestInitStraigthLaying(t *testing.T) {
	straigth := Init('-', point.Init(0, 1))
	res := straigth.Ways()
	expected := point.InitVecSet([]point.Vec{point.Init(0, 0), point.Init(0, 2)})
	if !expected.Eq(res) {
		t.Errorf("kolla närmare på ways av stående streck")
	}
}

func TestInitBendNortEast(t *testing.T) {
	bend := Init('L', point.Init(5, 5))
	res := bend.Ways()
	expected := point.InitVecSet([]point.Vec{point.Init(4, 5), point.Init(5, 6)})
	if !expected.Eq(res) {
		t.Errorf("titta närmare på L - biten")
	}
}

func TestInitBendNortWest(t *testing.T) {
	bend := Init('J', point.Init(5, 5))
	res := bend.Ways()
	expected := point.InitVecSet([]point.Vec{point.Init(5, 4), point.Init(4, 5)})
	if !expected.Eq(res) {
		t.Errorf("titta närmare på J - biten")
	}
}

func TestInitBendSoutWests(t *testing.T) {
	bend := Init('7', point.Init(5, 5))
	res := bend.Ways()
	expected := point.InitVecSet([]point.Vec{point.Init(5, 4), point.Init(6, 5)})
	if !expected.Eq(res) {
		t.Errorf("titta närmare på 7 - biten")
	}
}

func TestInitBendSoutEaasts(t *testing.T) {
	bend := Init('F', point.Init(5, 5))
	res := bend.Ways()
	expected := point.InitVecSet([]point.Vec{point.Init(6, 5), point.Init(5, 6)})
	if !expected.Eq(res) {
		t.Errorf("titta närmare på F - biten")
	}
}
