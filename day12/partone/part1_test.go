package partone

import (
	"day12/field"
	"day12/order"
	"testing"
)

// flytta dessa tester ner till field istället. De passar bättre där

func TestTrivialBaseCase(t *testing.T) {
	f := field.Init("#")
	o := order.Init([]int{1})
	res := NumberOfCombination(f, o)
	if res != 1 {
		t.Errorf("Något är brutalt fel")
	}
}

func TestSmpleBaseCase(t *testing.T) {
	f := field.Init("?.?")
	o := order.Init([]int{1})
	res := NumberOfCombination(f, o)
	if res != 2 {
		t.Errorf("nummber of combination returned wrong in the case ?.? 1")
	}
}

func TestRestriktion(t *testing.T) {
	f := field.Init("#.?")
	o := order.Init([]int{1})
	res := NumberOfCombination(f, o)
	if res != 1 {
		t.Errorf("dags att läsa instructionerna igen #.? 1 , ska vara 1 var  med var %d", res)
	}
}
