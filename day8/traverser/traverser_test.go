package traverser

import (
	"day8/graph"
	"testing"
)

func TestSimpleExample(t *testing.T) {
	inputs := []graph.NameLeftRight{
		{Name: "AAA", Left: "BBB", Right: "BBB"},
		{Name: "BBB", Left: "AAA", Right: "ZZZ"},
		{Name: "ZZZ", Left: "ZZZ", Right: "ZZZ"}}
	traveler := Init("LLR", inputs)
	res := traveler.NrOfSteps()
	if res != 6 {
		t.Errorf("the short example was not done correct")
	}
}
