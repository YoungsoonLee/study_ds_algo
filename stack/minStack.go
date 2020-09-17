package stack

import "math"

type MinStack struct {
	elementStack []int
	minimumStack []int
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(data int) {
	stack.elementStack = append(stack.elementStack, data)
	if len(stack.minimumStack) == 0 {
		stack.minimumStack = append(stack.minimumStack, data)
	} else {
		minimum := min(stack.minimumStack[len(stack.minimumStack)-1], data)
		stack.minimumStack = append(stack.minimumStack, minimum)
	}
}

func (stack *MinStack) Pop() int {
	if len(stack.elementStack) > 0 {
		poped := stack.elementStack[len(stack.elementStack)-1]
		stack.elementStack = stack.elementStack[:len(stack.elementStack)-1]
		stack.minimumStack = stack.minimumStack[:len(stack.minimumStack)-1]
		return poped
	} else {
		return math.MaxInt32
	}
}

func (stack *MinStack) Peek() int {
	if len(stack.elementStack) > 0 {
		return stack.elementStack[len(stack.elementStack)-1]
	} else {
		return 0
	}
}
