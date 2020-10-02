package main

import "fmt"

var total = 1

const N = 3 // n-ary

type NaryTreeNode struct {
	parent   *NaryTreeNode
	children []*NaryTreeNode
	depth    int
	data     int
}

func BuildNaryTree() *NaryTreeNode {
	var root NaryTreeNode
	root.addChildern()

	for _, child := range root.children {
		child.addChildern()
	}
	return &root
}

func (n *NaryTreeNode) addChildern() {
	for i := 0; i < N; i++ {
		newChild := &NaryTreeNode{
			parent: n,
			depth:  n.depth + 1,
		}
		n.children = append(n.children, newChild)
	}
}

func (n *NaryTreeNode) addChild() {
	newChild := &NaryTreeNode{
		parent: n,
		depth:  n.depth + 1,
	}
	n.children = append(n.children, newChild)
}

func (n *NaryTreeNode) walk() {
	n.visit()
	for _, child := range n.children {
		child.walk()
	}
}

func (n *NaryTreeNode) visit() {
	d := "L"
	for i := 0; i <= n.depth; i++ {
		d = d + "---------"
	}
	fmt.Printf("%s Visiting NaryTreeNNode with addr %p and parent %p total(%d)\n", d, n, n.parent, total)
	total = total + 1
}

func main() {
	fmt.Println("Building N-ary Tree")
	root := BuildNaryTree()
	fmt.Println("Walking n-ary tree")
	root.walk()
}
