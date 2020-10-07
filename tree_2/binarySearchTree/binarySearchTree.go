package main

import (
	"fmt"
)

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

func find(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	if root.data > data {
		return find(root.left, data)
	} else if root.data < data {
		return find(root.right, data)
	}

	return root

}

func findWithLoop(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	for root != nil {
		if root.data == data {
			return root
		} else if root.data > data {
			root = root.left
		} else {
			root = root.right
		}
	}

	return nil
}

func findMin(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	} else if root.left == nil {
		return root
	} else {
		return findMin(root.left)
	}
}

func findMinWithLoop(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

func findMax(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	} else if root.right == nil {
		return root
	} else {
		return findMax(root.right)
	}

}

func findMaxWithLoop(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}

	for root.right != nil {
		root = root.right
	}

	return root
}

func Insert(root *BSTNode, v int) *BSTNode {
	if root == nil {
		return &BSTNode{v, nil, nil}
	}
	if v < root.data {
		root.left = Insert(root.left, v)
	}
	root.right = Insert(root.right, v)
	return root
}

// !!!
func Delete(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	if data < root.data {
		root.left = Delete(root.left, data)
	} else if data > root.data {
		root.right = Delete(root.right, data)
	} else {
		if root.right == nil {
			return root.left
		}
		if root.left == nil {
			return root.right
		}
		t := root
		root = findMin(t.right)
		root.right = deleteMin(t.right)
		root.left = t.left
	}
	return root
}

func deleteMin(root *BSTNode) *BSTNode {
	if root.left == nil {
		return root.right
	}
	root.left = deleteMin(root.left)
	return root
}

func Walk(root *BSTNode, ch chan int) {
	if root == nil {
		return
	}
	Walk(root.left, ch)
	ch <- root.data
	Walk(root.right, ch)
}

func Walker(root *BSTNode) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(root, ch)
		close(ch)
	}()
	return ch
}

func Compare(t1, t2 *BSTNode) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func LCA(root *BSTNode, a, b int) *BSTNode {
	cur := root
	for {
		switch {
		case a < cur.data && b < cur.data:
			cur = cur.left
		case a > cur.data && b > cur.data:
			cur = cur.right
		default:
			return cur
		}
	}
	return root
}

func IsBST(root *BSTNode) bool {
	if root == nil {
		return true
	}

	if root.left != nil && root.left.data > root.data {
		return false
	}

	if root.right != nil && root.right.data < root.data {
		return false
	}

	if !IsBST(root.left) || !IsBST(root.right) {
		return false
	}

	return true
}

func BST2DLL(root *BSTNode) {
	if root == nil || (root.left == nil && root.right == nil) {
		return
	}

	BST2DLL(root.left)
	BST2DLL(root.right)
	currRight := root.right
	root.right = root.left
	root.left = nil
	for root.right != nil {
		root = root.right
	}
	root.right = currRight
}

func SortedArrayToBST(A []int) *BSTNode {
	if A == nil {
		return nil
	}

	return helper(A, 0, len(A)-1)
}

func helper(A []int, low int, high int) *BSTNode {
	if low > high {
		return nil
	}

	mid := low + (high-low)/2
	node := new(BSTNode)
	node.data = A[mid]
	node.left = helper(A, low, mid-1)
	node.right = helper(A, mid+1, high)
	return node
}

func SplitList(head *ListNode) (*ListNode, *ListNode, *ListNode) {
	if head.next == nil {
		return nil, head, nil
	}

	slowPointer := head
	fastPointer := head
	previousPointer := head

	for fastPointer != nil && fastPointer.next != nil {
		previousPointer = slowPointer
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next
	}
	previousPointer.next = nil
	return head, slowPointer, slowPointer.next
}

func SortedListToBST(head *ListNode) *BSTNode {
	if head == nil {
		return nil
	}

	left, middle, right := SplitList(head)

	var node BSTNode
	node.data = middle.data
	node.left = SortedListToBST(left)
	node.right = SortedListToBST(right)
	return &node
}

func kthSmallest(root *BSTNode, k int) *BSTNode {
	counter := 0
	return helper2(root, k, &counter)
}

func helper2(root *BSTNode, k int, counter *int) *BSTNode {
	if root == nil {
		return nil
	}

	left := helper2(root.left, k, counter)
	if left != nil {
		return left
	}
	*counter += 1
	if *counter == k {
		return root
	}
	return helper2(root.right, k, counter)
}

func RangePrinter(root *BSTNode, K1, K2 int) {
	if root == nil {
		return
	}

	if root.data >= K1 {
		RangePrinter(root.left, K1, K2)
	}
	if root.data >= K1 && root.data <= K2 {
		fmt.Print(" ", root.data)
	}
	if root.data <= K2 {
		RangePrinter(root.right, K1, K2)
	}
}

func RangePrinterWithQ(root *BSTNode, K1, K2 int) {
	if root == nil {
		return
	}
	var result [][]int
	queue := []*BSTNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.data >= K1 && node.data <= K2 {
				fmt.Print(" ", node.data)
			}
			if node.left != nil && node.data >= K1 {
				queue = append(queue, node.left)
			}
			if node.right != nil && node.data <= K2 {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
}

func CountTree(n int) int {
	if n <= 1 {
		return 1
	} else {
		sum := 0
		for root := 1; root <= n; root++ {
			left := CountTree(root - 1)
			right := CountTree(n - root)
			sum += left * right
		}
		return sum
	}
}
