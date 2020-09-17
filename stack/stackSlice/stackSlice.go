package stackSlice

import "errors"

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	return stack
}

func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity) // !!!!
}

func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) Resize() {
	if stack.IsFull() {
		stack.capacity *= 2
	} else {
		stack.capacity /= 2
	}
	target := make([]interface{}, stack.capacity)
	copy(target, stack.array[:stack.top+1])
	stack.array = target
}

func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		stack.Resize()
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *Stack) Pop() (interface{}, error) {
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

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := stack.array[stack.top]
	return temp, nil
}

func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}
