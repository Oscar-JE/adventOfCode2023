package workflow

import "day19/item"

type workNode interface {
	work(item.Item)
}

type rule struct {
	check item.Demand
	next  workNode
}

type checkNode struct {
	name  string
	rules []rule
}

func (c *checkNode) work(it item.Item) {
	for _, rule := range c.rules {
		if rule.check.Check(it) {
			rule.next.work(it)
		}
	}
}

type sumNode struct {
	name string
	sum  int
}

func (s *sumNode) work(it item.Item) {
	s.sum += it.Score()
}

type Workflow struct {
	startNode  checkNode
	acceptNode sumNode
	rejectNode sumNode
}

type WorkInformation struct {
	nodeId string
}

type NodeInfo struct {
	Id    string
	Rules []RuleRep
}

type RuleRep struct {
	Demand        item.Demand
	DestinationId string
}

func Init(nodes []NodeInfo) Workflow {
	// antar att första noden ska in här
	// skapa upp med nill varden och pupulera sedan
	return Workflow{}
}
