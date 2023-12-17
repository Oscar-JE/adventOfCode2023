package graph

import (
	"testing"
)

func TestCreateGraph(t *testing.T) {
	inputs := []NameLeftRight{
		{Name: "AAA", Left: "BBB", Right: "BBB"},
		{Name: "BBB", Left: "AAA", Right: "ZZZ"},
		{Name: "ZZZ", Left: "ZZZ", Right: "ZZZ"}}

	root := Init(inputs)
	b := *root.left
	z := *b.right
	if z.name != "ZZZ" {
		t.Errorf("Check how the graph is setup")
	}
}

func TestInitAndFindStartingNodes(t *testing.T) {
	inputs := []NameLeftRight{
		{Name: "11A", Left: "11B", Right: "XXX"},
		{Name: "11B", Left: "XXX", Right: "11Z"},
		{Name: "11Z", Left: "11B", Right: "XXX"},
		{Name: "22A", Left: "22B", Right: "XXX"},
		{Name: "22B", Left: "22C", Right: "22C"},
		{Name: "22C", Left: "22B", Right: "22B"},
		{Name: "XXX", Left: "XXX", Right: "XXX"}}
	startNodes := InitFindStartingNodes(inputs)
	if len(startNodes) != 2 || startNodes[0].GetName() != "11A" || startNodes[1].GetName() != "22A" {
		t.Errorf("inspect InitFindStartingNodes")
	}
}
