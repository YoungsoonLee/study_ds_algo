package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (t *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if t.Val > v {
		if t.Left == nil {
			t.Left = &BinaryTreeNode{Val: v}
			return t.Left
		} else {
			return t.Left.AddNode(v)
		}
	} else {
		if t.Right == nil {
			t.Right = &BinaryTreeNode{Val: v}
			return t.Right
		} else {
			return t.Right.AddNode(v)
		}
	}
}

func (t *BinaryTree) Search(v int) (bool, int) {
	return t.Root.Search(v, 1)
}

// DFS
func (n *BinaryTreeNode) Search(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.Left != nil {
			return n.Left.Search(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.Right != nil {
			return n.Right.Search(v, cnt+1)
		}
		return false, cnt
	}
}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

// BFS는 queue를 사용 해서...
func (t *BinaryTree) PrintBFS() {
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})

	currentDepth := 0

	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			//fmt.Println()
			currentDepth = first.depth
		}

		fmt.Print(first.node.Val, " ")

		if first.node.Left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Left})
		}

		if first.node.Right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Right})
		}

	}
}

// DFS는 stack을 사용 해서...
func (t *BinaryTree) PrintDFS() {
	s := []depthNode{}
	s = append(s, depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(s) > 0 {
		var first depthNode
		first, s = s[len(s)-1], s[0:len(s)-1]

		if first.depth != currentDepth {
			//fmt.Println()
			currentDepth = first.depth
		}

		fmt.Print(first.node.Val, " ")

		if first.node.Right != nil {
			s = append(s, depthNode{depth: currentDepth + 1, node: first.node.Right})
		}

		if first.node.Left != nil {
			s = append(s, depthNode{depth: currentDepth + 1, node: first.node.Left})
		}

	}
}

func (b *BinaryTreeNode) PrintDFS_by_recursive() {
	if b == nil {
		return
	}

	fmt.Printf("%d ", b.Val)

	if b.Left != nil {
		b.Left.PrintDFS_by_recursive()
	}

	if b.Right != nil {
		b.Right.PrintDFS_by_recursive()
	}
}
