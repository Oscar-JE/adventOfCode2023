package set

import "testing"

func TestAdd(t *testing.T) {
	disjunctIntervals := []Interval{{0, 10}, {20, 30}, {40, 50}, {60, 70}, {80, 90}}
	mySet := Set{disjunctIntervals: disjunctIntervals}
	disjunctExpected := []Interval{{0, 10}, {20, 70}, {80, 90}}
	expected := Set{disjunctExpected}
	inter := Interval{25, 65}
	mySet.add(inter)
	if !mySet.Eq(expected) {
		t.Errorf("\n got %v  \n Expected %v", mySet, expected)
	}
}

func TestAddToEmpty(t *testing.T) {
	disjunctIntervals := []Interval{}
	mySet := Set{disjunctIntervals: disjunctIntervals}
	inter := Interval{0, 10}

	disjunctExpected := []Interval{{0, 10}}
	expected := Set{disjunctExpected}

	mySet.add(inter)
	if !mySet.Eq(expected) {
		t.Errorf("\n got %v  \n Expected %v", mySet, expected)
	}
}

func TestAddLast(t *testing.T) {
	disjunctIntervals := []Interval{{0, 10}}
	mySet := Set{disjunctIntervals: disjunctIntervals}
	inter := Interval{20, 30}

	disjunctExpected := []Interval{{0, 10}, {20, 30}}
	expected := Set{disjunctExpected}

	mySet.add(inter)
	if !mySet.Eq(expected) {
		t.Errorf("\n got %v  \n Expected %v", mySet, expected)
	}
}

func TestAddFirst(t *testing.T) {
	disjunctIntervals := []Interval{{20, 30}}
	mySet := Set{disjunctIntervals: disjunctIntervals}
	inter := Interval{0, 10}

	disjunctExpected := []Interval{{0, 10}, {20, 30}}
	expected := Set{disjunctExpected}

	mySet.add(inter)
	if !mySet.Eq(expected) {
		t.Errorf("\n got %v  \n Expected %v", mySet, expected)
	}
}
