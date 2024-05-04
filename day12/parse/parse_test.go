package parse

import (
	"day12/field"
	"day12/order"
	"day12/springs"
	"testing"
)

func TestParseLine(t *testing.T) {
	inputLine := "? 4,4"
	f, positions := Line(inputLine)
	expectedF := field.InitFromSprings([]springs.Condition{springs.Unknown()})
	expectedOrder := order.Init([]int{4, 4})
	if !f.Eq(expectedF) || !positions.Eq(expectedOrder) {
		t.Errorf("parse out of controll")
	}

}
