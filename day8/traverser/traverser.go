package traverser

import (
	"day8/graph"
	"day8/policy"
)

func NrOfSteps(pol string, graphInfo []graph.NameLeftRight) int {
	poli := policy.Init(pol)
	node := graph.Init(graphInfo)
	nrOfsteps := 0
	for nrOfsteps < 10000000 {
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
	for nrOfsteps < 1000000000000 {
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
