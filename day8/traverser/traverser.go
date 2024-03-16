package traverser

import (
	"day8/congruence"
	"day8/graph"
	"day8/policy"
)

func FindLoops(pol string, graphInfo []graph.NameLeftRight) []congruence.TailedLoopMultipleEndPoints {
	poli := policy.Init(pol)
	nodes := graph.InitFindStartingNodes(graphInfo)
	lopes := []congruence.TailedLoopMultipleEndPoints{}
	for i := range nodes {
		poli.Recett()
		lopes = append(lopes, findLoop(poli, nodes[i]))
	}
	return lopes
}

type checkpoint struct {
	decisionIndex int
	nodeName      string
}

func eq(h1 []checkpoint, h2 []checkpoint) bool {
	if len(h1) != len(h2) {
		return false
	}
	equal := true
	for i := range h1 {
		equal = equal && h1[i] == h2[i]
	}
	return equal
}

func createLoopFromHistory(history []checkpoint) congruence.TailedLoopMultipleEndPoints {
	return congruence.InitTailedLoop(0, 0, []int{})
}

func findLoop(pol policy.Policy, node graph.Node) congruence.TailedLoopMultipleEndPoints {
	findHistory(pol, node)
	return congruence.InitTailedLoop(0, 0, []int{0})
}

func findHistory(pol policy.Policy, node graph.Node) []checkpoint {
	history := []checkpoint{}
	foundInHistory := false
	currentNode := node
	for !foundInHistory {
		decision, index := pol.GetNextAndindex()
		c := checkpoint{index, currentNode.GetName()}
		currentNode = currentNode.Go(decision)
		foundInHistory = search(history, c)
		history = append(history, c)
	}
	return history
}

func search(history []checkpoint, c checkpoint) bool {
	for _, el := range history {
		if el == c {
			return true
		}
	}
	return false
}

func NrOfSteps(pol string, graphInfo []graph.NameLeftRight) int {
	poli := policy.Init(pol)
	node := graph.Init(graphInfo)
	nrOfsteps := 0
	for nrOfsteps < 1000 {
		if node.GetName() == "ZZZ" {
			return nrOfsteps
		}
		decision := poli.GetNext()
		node = node.Go(decision)
		nrOfsteps++
	}
	return nrOfsteps
}

func NrOfMutualSteps(polStr string, graphInfo []graph.NameLeftRight) int {
	poli := policy.Init(polStr)
	nodes := graph.InitFindStartingNodes(graphInfo)
	nrOfsteps := 0
	for nrOfsteps < 1000000 {
		if reachedEnd(nodes) {
			return nrOfsteps
		}
		decision := poli.GetNext()
		for i := range nodes {
			nodes[i] = nodes[i].Go(decision)
		}
		nrOfsteps++
	}
	return nrOfsteps
}

func reachedEnd(nodes []graph.Node) bool {
	condition := true
	for _, node := range nodes {
		condition = condition && node.GetName()[len(node.GetName())-1] == 'Z'
	}
	return condition
}