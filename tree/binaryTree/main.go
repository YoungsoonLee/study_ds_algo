package main

import (
	"fmt"
)

func main() {
	tree := NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.PrintBFS()

	fmt.Println()
	fmt.Println("--------------------")

	tree.PrintDFS()
	fmt.Println()
	fmt.Println("--------------------")

	if found, cnt := tree.Search(6); found {
		fmt.Println("found 6, cnt: ", cnt)
	} else {
		fmt.Println("not found 6, cnt: ", cnt)
	}

	if found, cnt := tree.Search(11); found {
		fmt.Println("found 11, cnt: ", cnt)
	} else {
		fmt.Println("not found 11, cnt: ", cnt)
	}

	fmt.Println("--------------------")

	tv2 := NewBTree(5)
	tv2.AddNode(3)
	tv2.AddNode(2)
	tv2.AddNode(4)
	tv2.AddNode(8)
	tv2.AddNode(7)
	tv2.AddNode(6)
	tv2.AddNode(10)
	tv2.AddNode(9)

	fmt.Println()
	fmt.Println("--------------------")
	tv2.PrintDFS()
	fmt.Println()
	fmt.Println("--------------------")
	tv2.PrintBFS()
	//https://qa-closers.naddicjapan.com/dmm.do
	//https://closers.naddicjapan.com/index_.do
}
