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
