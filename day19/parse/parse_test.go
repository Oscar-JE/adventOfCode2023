package parse

import (
	"day19/item"
	"day19/workflow"
	"fmt"
	"testing"
)

func TestParseItem(t *testing.T) {
	strRep := "{x=1,m=2,a=3,s=4}"
	expected := item.Init(1, 2, 3, 4)
	res := parseItem(strRep)
	if res != expected {
		t.Errorf("parse item does not work")
	}
}

func TestParseRule(t *testing.T) {
	strRep := "a<2006:qkq"
	expected := workflow.RuleRep{Demand: item.LessAerodynamic{2006}, DestinationId: "qkq"}
	res := parseRule(strRep)
	if expected != res {
		t.Errorf("Parse demand needs more attention")
	}
}

func TestParseWorkflow(t *testing.T) {
	strRep := "px{m>2090:A,rfg}"
	res := parseNode(strRep)
	fmt.Println(res)
}
