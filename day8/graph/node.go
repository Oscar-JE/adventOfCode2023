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
	panic("no node with that name was found")
}

func (n Node) Go(dir direction.Direction) Node {
	if dir == direction.LEFT {
		return *n.left
	} else {
		return *n.right
	}
}
