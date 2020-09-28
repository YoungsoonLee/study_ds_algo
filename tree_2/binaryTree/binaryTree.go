package main

import (
	"fmt"
	"math"
	"math/rand"
	// stack "github.com/YoungsoonLee/study-ds-algo/stack"
)

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// Preorder: DLR
// DFS
func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d", root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

func PreOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}

	ch <- root.data
	PreOrderWalk(root.left, ch)
	PreOrderWalk(root.right, ch)
}

func PreOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		go PreOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

func NewBinaryTree(n, k int) *BinaryTreeNode {
	var root *BinaryTreeNode
	for _, v := range rand.Perm(n) {
		root = insert(root, (1+v)*k)
	}
	return root
}

func insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{v, nil, nil}
	}
	if v < root.data {
		root.left = insert(root.left, v)
		return root
	}
	root.right = insert(root.right, v)
	return root
}

func InOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Printf("%d", root.data)
	InOrder(root.right)
}

func InOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}
	InOrderWalk(root.left, ch)
	ch <- root.data
	InOrderWalk(root.right, ch)
}

func InOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		go InOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

func PostOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d", root.data)
}

func PostOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}
	PostOrderWalk(root.left, ch)
	PostOrderWalk(root.right, ch)
	ch <- root.data
}

func PostOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		PostOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

// using queue
func LevelOrder(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
	return result
}

func findMax(root *BinaryTreeNode) int {
	max := math.MinInt32
	if root != nil {
		root_val := root.data
		left := findMax(root.left)
		right := findMax(root.right)

		if left > right {
			max = left
		} else {
			max = right
		}

		if root_val > max {
			max = root_val
		}
	}
	return max
}

func findMaxWithQueue(root *BinaryTreeNode) int {
	max := math.MinInt32
	if root == nil {
		return max
	}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			if node.data > max {
				max = node.data
			}
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return max
}

func find(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		return root
	} else {
		if data == root.data {
			return root
		} else {
			temp := find(root.left, data)
			if temp != nil {
				return temp
			} else {
				return find(root.right, data)
			}
		}
	}
}

func findWithQueue(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		return root
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			if node.data == data {
				return node
			}
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return nil
}

func Insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{v, nil, nil}
	if root == nil {
		return newNode
	}
	if root.left == nil {
		root.left = insert(root.left, v)
	} else if root.right == nil {
		root.right = insert(root.right, v)
	} else {
		randomize := rand.Intn(1)
		if randomize == 0 {
			root.left = insert(root.left, v)
		} else {
			root.right = insert(root.right, v)
		}
	}
	return root
}

func InsertWithQueue(root *BinaryTreeNode, v int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{v, nil, nil}
	if root == nil {
		return newNode
	}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			} else {
				node.left = newNode
				return root
			}
			if node.right != nil {
				queue = append(queue, node.right)
			} else {
				node.right = newNode
				return root
			}

		}
	}
	return root
}

func Size(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	} else {
		return 1 + Size(root.left) + Size(root.right)
	}
}

func SizeWithQueue(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			result++
			level = append(level, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return result
}

/*
func LevelOrderBottomUp(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*BinaryTreeNode{root}
	s := stack.NewStackSlice(1)

	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.right != nil {
				queue = append(queue, node.right)
			}
			if node.left != nil {
				queue = append(queue, node.left)
			}
		}
		s.Push(level)
	}

	for !s.IsEmpty() {
		result = append(result, s.Pop().([]int))
	}
	return result
}
*/

func DeleteTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	root.left = DeleteTree(root.left)
	root.right = DeleteTree(root.right)

	root = nil
	return root
}

func Height(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	} else {
		leftHeight := Height(root.left)
		rightHeight := Height(root.right)
		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
}

func HeightWithQueue(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		count++
	}
	return count
}

func Deepest(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	var node *BinaryTreeNode
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		//var level []int
		for i := 0; i < qlen; i++ {
			node = queue[0]
			//level = append(level, node.data)
			queue = queue[1:]

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return node
}

func LeavesCount(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		//var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			//level = append(level, node.data)
			queue = queue[1:]
			if node.left == nil && node.right == nil {
				count++
			}
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return count
}

func FullNodesCountWithQ(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil && node.right != nil {
				count++
			}
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return count
}

// recursion
func FullNodesCount(root *BinaryTreeNode) int {

	count := 0
	//left := 0
	//right := 0
	if root.left != nil && root.right != nil {
		count++
	}

	if root.left != nil {
		count += FullNodesCount(root.left)
	}

	if root.right != nil {
		count += FullNodesCount(root.right)
	}

	return count
}

func HalfNodesCountWithQ(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil && node.right == nil {
				count++
			} else if node.left == nil && node.right != nil {
				count++
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return count
}

func HalfNodesCount(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	//left := 0
	//right := 0

	if (root.left != nil && root.right == nil) || (root.left == nil && root.right != nil) {
		count++
	}
	if root.left != nil {
		count += HalfNodesCount(root.left)
	}
	if root.right != nil {
		count += HalfNodesCount(root.right)
	}
	return count
}

func CompareStructures(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return CompareStructures(root1.left, root2.left) && CompareStructures(root1.right, root2.right)
}

func DiameterOfBinaryTree(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	var diameter int
	Diameter(root, &diameter)
	return diameter
}

func Diameter(root *BinaryTreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	leftDepth := Diameter(root.left, diameter)
	rightDepth := Diameter(root.right, diameter)

	if leftDepth+rightDepth > *diameter {
		*diameter = leftDepth + rightDepth
	}
	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxLevelSum(root *BinaryTreeNode) (ele []int, maxSum, level int) {
	ele, maxSum, level = []int{}, math.MinInt32, 0
}

func main() {
	t1 := NewBinaryTree(10, 1)
	PreOrder(t1)
	fmt.Println()
	/*
		c := PreOrderWalker(t1)
		for {
			v, ok := <-c
			if !ok {
				break
			}
			fmt.Printf("%d", v)
		}
	*/

	fmt.Println("findMax: ", findMax(t1))

	fmt.Println("Deepest: ", Deepest(t1))

	fmt.Println("LeavesCount: ", LeavesCount(t1))

	fmt.Println("FullNodesCount: ", FullNodesCount(t1))

	fmt.Println("FullNodesCountWithQ: ", FullNodesCountWithQ(t1))

	fmt.Println("HalfNodesCountWithQ: ", HalfNodesCountWithQ(t1))

	fmt.Println("HalfNodesCount: ", HalfNodesCount(t1))
}
