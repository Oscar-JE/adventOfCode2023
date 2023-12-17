package traverser

import (
	"day8/graph"
	"day8/policy"
)

type Traverser struct { //det här är en struct vi egentligen inte behöver kunde varit en function
	pol  policy.Policy
	node graph.Node
}

func Init(pol string, graphInfo []graph.NameLeftRight) Traverser {
	poli := policy.Init(pol)
	root := graph.Init(graphInfo)
	return Traverser{pol: poli, node: root}
}

func (t Traverser) NrOfSteps() int {
	nrOfsteps := 0
	for nrOfsteps < 10000000 {
		if t.node.GetName() == "ZZZ" {
			return nrOfsteps
		}
		decision := t.pol.GetNext()
		t.node = t.node.Go(decision)
		nrOfsteps++
	}
	return nrOfsteps
}
