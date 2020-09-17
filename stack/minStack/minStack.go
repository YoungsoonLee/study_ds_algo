package minStack

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

func (stack *MinStack) Size() int {
	return len(stack.elementStack)
}

func (stack *MinStack) GetMin() int {
	if len(stack.minimumStack) > 0 {
		return stack.minimumStack[len(stack.minimumStack)-1]
	} else {
		return 0
	}
}

func (stack *MinStack) IsEmpty() bool {
	return len(stack.elementStack) == 0
}

func (stack *MinStack) Clear() {
	stack.elementStack = nil
	stack.minimumStack = nil
}

func reverse(str string) string {
	result := []rune(str)
	var beg int
	end := len(result) - 1

	for beg < end {
		result[beg], result[end] = result[end], result[beg]
		beg = beg + 1
		end = end - 1
	}
	return string(result)
}

func isPalindrome(testString string) bool {
	if reverse(testString) == testString {
		return true
	}
	return false
}

func isPalindrome2(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] { // !!!
			return false
		}
	}
	return true
}
