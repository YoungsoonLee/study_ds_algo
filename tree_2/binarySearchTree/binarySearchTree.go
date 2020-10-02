package main

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
