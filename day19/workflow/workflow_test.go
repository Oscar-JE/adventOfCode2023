package workflow

import (
	"day19/item"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) { // kan väl räcka med den nivån
	nodes := []NodeInfo{}
	nodes = append(nodes, NodeInfo{Id: "in", Rules: []RuleRep{{Demand: item.LessMusical{Musical: 4}, DestinationId: "A"}, {Demand: item.NoDemand{}, DestinationId: "R"}}})
	nodes = append(nodes, NodeInfo{Id: "A", Rules: []RuleRep{}})
	nodes = append(nodes, NodeInfo{Id: "R", Rules: []RuleRep{}})
	workFlow := Init(nodes)
	fmt.Println(workFlow)
}
