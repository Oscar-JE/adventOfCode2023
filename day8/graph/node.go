package graph

import "day8/direction"

type Node struct {
	name  string
	left  *Node
	right *Node
}

func (n Node) GetName() string {
	return n.name
}

type NameLeftRight struct {
	Name  string
	Left  string
	Right string
}

func Init(inputs []NameLeftRight) Node {
	nodes := []Node{}
	for _, nameLR := range inputs {
		nodes = append(nodes, Node{nameLR.Name, nil, nil})
	}
	for i, nameRL := range inputs {
		nodes[i].left = findNodeWithName(nodes, nameRL.Left)
		nodes[i].right = findNodeWithName(nodes, nameRL.Right)
	}
	ref := findNodeWithName(nodes, "AAA")
	return *ref
}

func findNodeWithName(nodes []Node, name string) *Node {
	for i, node := range nodes {
		if node.name == name {
			return &nodes[i]
		}
	}
	panic("no node with that name was found : " + name)
}

func (n Node) Go(dir direction.Direction) Node {
	if dir == direction.LEFT {
		return *n.left
	} else {
		return *n.right
	}
}

func InitFindStartingNodes(inputs []NameLeftRight) []Node {
	nodes := []Node{}
	for _, nameLR := range inputs {
		nodes = append(nodes, Node{nameLR.Name, nil, nil})
	}
	for i, nameRL := range inputs {
		nodes[i].left = findNodeWithName(nodes, nameRL.Left)
		nodes[i].right = findNodeWithName(nodes, nameRL.Right)
	}
	startingNodes := findStartingNodes(nodes)
	return startingNodes
}

func findStartingNodes(graph []Node) []Node {
	startingNodes := []Node{}
	for _, node := range graph {
		if node.name[len(node.name)-1] == 'A' {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}
