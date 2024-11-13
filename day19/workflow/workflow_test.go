package workflow

import (
	"day19/item"
	"fmt"
	"testing"
)

func TestSimpleWorkflowNoAccepted(t *testing.T) {
	nodes := []NodeInfo{}
	nodes = append(nodes, NodeInfo{Id: "in", Rules: []RuleRep{{Demand: item.LessMusical{Musical: 1}, DestinationId: "A"}, {Demand: item.NoDemand{}, DestinationId: "R"}}})
	nodes = append(nodes, NodeInfo{Id: "A", Rules: []RuleRep{}})
	nodes = append(nodes, NodeInfo{Id: "R", Rules: []RuleRep{}})
	workFlow := Init(nodes)
	res := workFlow.FindNumberOfPossibleAccept()
	if res != 0 {
		t.Errorf("impossible demans should result in zero combination")
	}
	fmt.Println(workFlow)
}

func TestSimpleWorkflowOneAccepted(t *testing.T) {
	nodes := []NodeInfo{}
	nodes = append(nodes, NodeInfo{Id: "in", Rules: []RuleRep{{Demand: item.LessMusical{Musical: 2}, DestinationId: "A"}, {Demand: item.NoDemand{}, DestinationId: "R"}}})
	nodes = append(nodes, NodeInfo{Id: "A", Rules: []RuleRep{}})
	nodes = append(nodes, NodeInfo{Id: "R", Rules: []RuleRep{}})
	workFlow := Init(nodes)
	res := workFlow.FindNumberOfPossibleAccept()
	expected := 4000 * 4000 * 4000 * 1
	if res != expected {
		t.Errorf("expected: %d but gott %d", expected, res)
	}
	fmt.Println(workFlow)
}

/*
func TestTwoLevels(t *testing.T) {
	nodes := []NodeInfo{}
	nodes = append(nodes, NodeInfo{Id: "in", Rules: []RuleRep{{Demand: item.LessMusical{Musical: 5}, DestinationId: "A"}, {Demand: item.NoDemand{}, DestinationId: "R"}}})
	nodes = append(nodes, NodeInfo{Id: "m1", Rules: []RuleRep{{Demand: item.MoreMusical{6}, DestinationId: "R"},De}})
	nodes = append(nodes, NodeInfo{Id: "A", Rules: []RuleRep{}})
	nodes = append(nodes, NodeInfo{Id: "R", Rules: []RuleRep{}})
	workFlow := Init(nodes)
}*/
