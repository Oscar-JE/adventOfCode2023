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
			return
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
	startNode  *checkNode
	acceptNode *sumNode
	rejectNode *sumNode
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

func (w Workflow) ProcessItems(items []item.Item) int {
	for _, item := range items {
		w.process(item)
	}
	return w.acceptNode.sum
}

func (w *Workflow) process(item item.Item) {
	activeNode := w.startNode
	activeNode.work(item)
}

func Init(nodes []NodeInfo) Workflow {
	workNodes := []*checkNode{}
	accept := sumNode{name: "A", sum: 0}
	reject := sumNode{name: "R", sum: 0}
	for _, nodeInf := range nodes {
		loopNode := checkNode{nodeInf.Id, []rule{}}
		workNodes = append(workNodes, &loopNode)
	}

	for index, nodeInf := range nodes {
		workNodes[index].rules = append(workNodes[index].rules, createRule(nodeInf.Rules, workNodes, &accept, &reject)...)
	}
	startnode := findStartNode(workNodes)
	workflow := Workflow{startNode: startnode, rejectNode: &reject, acceptNode: &accept}
	return workflow
}

func createRule(ruleReps []RuleRep, workNodes []*checkNode, accept *sumNode, reject *sumNode) []rule {
	retRules := []rule{}
	for _, ruleRep := range ruleReps {
		nextNodeId := ruleRep.DestinationId
		nextNode := findNextNode(nextNodeId, workNodes, accept, reject)
		retRules = append(retRules, rule{check: ruleRep.Demand, next: nextNode})
	}
	return retRules
}

func findNextNode(nextNodeId string, workNodes []*checkNode, accept *sumNode, reject *sumNode) workNode {
	if nextNodeId == "A" {
		return accept
	}
	if nextNodeId == "R" {
		return reject
	}
	for _, workNode := range workNodes {
		if nextNodeId == workNode.name {
			return workNode
		}
	}
	panic("No match for id was found")
}

func findStartNode(workNodes []*checkNode) *checkNode {
	for _, node := range workNodes {
		if node.name == "in" {
			return node
		}
	}
	panic("input node not present")
}