package traverser

import (
	"day8/congruence"
	"day8/graph"
	"day8/policy"
	"strings"
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
	reentry_index := 0 // failar om loopens längd är under 3
	last_element := history[len(history)-1]
	for i := len(history) - 2; i >= 0; i-- {
		if history[i] == last_element {
			reentry_index = i
			break
		}
	}
	loopLen := (len(history) - 1) - reentry_index // längd utan upprepningspunkten minus där vi hoppade in
	indixies_of_winingpoints := []int{}
	for i := 0; i < len(history)-1; i++ {
		if strings.Contains(history[i].nodeName, "Z") {
			indixies_of_winingpoints = append(indixies_of_winingpoints, i)
		}
	}
	return congruence.InitTailedLoop(reentry_index, loopLen, indixies_of_winingpoints)
}

func findLoop(pol policy.Policy, node graph.Node) congruence.TailedLoopMultipleEndPoints {
	history := findHistory(pol, node)
	return createLoopFromHistory(history)
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
