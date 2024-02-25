package traverser

import (
	"day8/congruence"
	"day8/graph"
	"day8/policy"
	"testing"
)

func simple1() []graph.NameLeftRight {
	return []graph.NameLeftRight{
		{Name: "AAA", Left: "BBB", Right: "BBB"},
		{Name: "BBB", Left: "AAA", Right: "ZZZ"},
		{Name: "ZZZ", Left: "ZZZ", Right: "ZZZ"}}
}

func TestFinndLoop(t *testing.T) {
	example := graph.Init(simple1())
	pol := policy.Init("LLR")
	loop := findLoop(pol, example)
	// AAA 0 , BBB 1, AAA 2 , BBB 0, AAA 1, BBB 2, ZZZ 0, ZZZ 1 , ZZZ 2  , ZZZ 0 hittatloop
	winingPoints := []int{6, 7, 8}
	tailLength := 6
	loopLength := 3
	expectedLoop := congruence.InitTailedLoop(tailLength, loopLength, winingPoints)
	if !expectedLoop.Eq(loop) {
		t.Errorf("unexpected loop")
	}
}

func simple2() []graph.NameLeftRight {
	return []graph.NameLeftRight{
		{Name: "11A", Left: "11B", Right: "XXX"},
		{Name: "11B", Left: "XXX", Right: "11Z"},
		{Name: "11Z", Left: "11B", Right: "XXX"},
		{Name: "22A", Left: "22B", Right: "XXX"},
		{Name: "22B", Left: "22C", Right: "22C"},
		{Name: "22C", Left: "22Z", Right: "22Z"},
		{Name: "22Z", Left: "22B", Right: "22B"},
		{Name: "XXX", Left: "XXX", Right: "XXX"}}
}

func TestSimpleExample(t *testing.T) {
	inputs := simple1()
	res := NrOfSteps("LLR", inputs)
	if res != 6 {
		t.Errorf("the short example was not done correct")
	}
}

func TestSimpleExamplePart2(t *testing.T) {
	inputs := simple2()
	res := NrOfMutualSteps("LR", inputs)
	if res != 6 {
		t.Errorf("Fasulty behaviour in Nr of mutual steps")
	}
}
