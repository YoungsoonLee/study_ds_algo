package main

import "errors"

//type GraphType string

type AdjacencyList struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjList   []*Node
}

type Node struct {
	Next   *Node
	Weight int
	Key    int
}

func (node Node) AddNode(value int) *Node {
	n := node.Next
	if n == nil {
		newNode := &Node{Next: &Node{}, Key: value}
		return newNode
	}
	nd := n.AddNode(value)
	node.Next = nd
	return &node
}

func (node Node) AddNodeWithWeight(value, weight int) *Node {
	n := node.Next
	if n == nil {
		newNode := &Node{Next: &Node{}, Key: value, Weight: weight}
		return newNode
	}
	nd := n.AddNodeWithWeight(value, weight)
	node.Next = nd
	return &node
}

func (node Node) FindNextNode(key int) (*Node, error) {
	n := node
	if n == (Node{}) {
		return &Node{}, errors.New("Node not found")
	}
	if n.Key == key {
		return &n, nil
	}
	nd := n.Next
	return nd.FindNextNode(key)
}
