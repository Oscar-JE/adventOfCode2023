package integration

import (
	"testing"
)

func TestTrivialCase(t *testing.T) {
	cols := []int{}
	res := RowIntegration(cols)
	expected := 0
	if res != expected {
		t.Errorf("Trivial case should be %d , but was %d", expected, res)
	}
}

func TestSinglularCase(t *testing.T) {
	borders := []int{5}
	res := RowIntegration(borders)
	expected := 1
	if res != expected {
		t.Errorf("Singular case should return %d but was %d", expected, res)
	}
}

func TestSimpleCase(t *testing.T) {
	boarders := []int{1, 2}
	res := RowIntegration(boarders)
	expected := 2
	if res != expected {
		t.Errorf("should be %d but was %d ", expected, res)
	}
}

func TestStandardCase(t *testing.T) {
	boarders := []int{1, 3}
	res := RowIntegration(boarders)
	expected := 3
	if res != expected {
		t.Errorf("should be %d but was %d ", expected, res)
	}
}
