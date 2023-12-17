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
	res := NrOfSteps("LLR", inputs)
	if res != 6 {
		t.Errorf("the short example was not done correct")
	}
}

func TestSimpleExamplePart2(t *testing.T) {
	inputs := []graph.NameLeftRight{
		{Name: "11A", Left: "11B", Right: "XXX"},
		{Name: "11B", Left: "XXX", Right: "11Z"},
		{Name: "11Z", Left: "11B", Right: "XXX"},
		{Name: "22A", Left: "22B", Right: "XXX"},
		{Name: "22B", Left: "22C", Right: "22C"},
		{Name: "22C", Left: "22Z", Right: "22Z"},
		{Name: "22Z", Left: "22B", Right: "22B"},
		{Name: "XXX", Left: "XXX", Right: "XXX"}}
	res := NrOfMutualSteps("LR", inputs)
	if res != 6 {
		t.Errorf("Fasulty behaviour in Nr of mutual steps")
	}
}
