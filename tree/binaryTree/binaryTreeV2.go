package main

import "fmt"

type BTree struct {
	Val   int
	Left  *BTree
	Right *BTree
}

func NewBTree(v int) *BTree {
	tree := &BTree{Val: v}
	return tree
}

func (b *BTree) AddNode(v int) *BTree {

	if b.Val > v {
		// left
		if b.Left == nil {
			b.Left = NewBTree(v)
			return b.Left
		} else {
			return b.Left.AddNode(v)
		}
	} else {
		// right
		if b.Right == nil {
			b.Right = NewBTree(v)
			return b.Right
		} else {
			return b.Right.AddNode(v)
		}
	}

}

func (b *BTree) PrintDFS() {
	if b == nil {
		return
	}

	fmt.Printf("%d ", b.Val)

	if b.Left != nil {
		b.Left.PrintDFS()
	}

	if b.Right != nil {
		b.Right.PrintDFS()
	}
}

// BFS is use queue !!!
func (b *BTree) PrintBFS() {
	q := []*BTree{}
	q = append(q, b)

	//fmt.Printf("%+v", q)
	fmt.Printf("init count: %d\n", len(q))

	for len(q) > 0 {

		var first *BTree // !!!
		first, q = q[0], q[1:]

		fmt.Printf("%d ", first.Val)

		if first.Left != nil {
			q = append(q, first.Left)
		}

		if first.Right != nil {
			q = append(q, first.Right)
		}

	}

}
