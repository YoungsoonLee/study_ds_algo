package models

import "errors"

type StackSlice struct {
	top      int
	capacity uint
	array    []interface{}
}

func (stack *StackSlice) Init(capacity uint) *StackSlice {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	return stack
}

func NewStackSlice(capacity uint) *StackSlice {
	return new(StackSlice).Init(capacity) // !!!!
}

func (stack *StackSlice) Size() uint {
	return uint(stack.top + 1)
}

func (stack *StackSlice) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

func (stack *StackSlice) IsEmpty() bool {
	return stack.top == -1
}

func (stack *StackSlice) Resize() {
	if stack.IsFull() {
		stack.capacity *= 2
	} else {
		stack.capacity /= 2
	}
	target := make([]interface{}, stack.capacity)
	copy(target, stack.array[:stack.top+1])
	stack.array = target
}

func (stack *StackSlice) Push(data interface{}) error {
	if stack.IsFull() {
		stack.Resize()
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *StackSlice) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.array[stack.top]
	stack.top--
	if stack.Size() < stack.capacity/2 {
		stack.Resize()
	}
	return temp, nil
}

func (stack *StackSlice) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.array[stack.top]
	return temp, nil
}

func (stack *StackSlice) Drain() {
	stack.array = nil
	stack.top = -1
}
