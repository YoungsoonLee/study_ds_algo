package models

import "errors"

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

func (s *Stack) Init(capacity uint) *Stack {
	s.top = -1
	s.capacity = capacity
	s.array = make([]interface{}, capacity)
	return s
}

func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity) // !!!!!
}

func (s *Stack) IsFull() bool {
	return s.top == int(s.capacity)-1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Size() uint {
	return uint(s.top + 1)
}

func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		return errors.New("Stack is full")
	}
	s.top++
	s.array[s.top] = data
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	temp := s.array[s.top]
	s.top--
	return temp, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	temp := s.array[s.top]
	return temp, nil
}

func (s *Stack) Drain() {
	s.array = nil
	s.top = -1
}

type pair struct {
	open  rune
	close rune
}

var pairs = []pair{
	{'(', ')'},
	{'[', ']'},
	{'{', '}'},
}

func isValid(s string) bool {
	stack := NewStack(1)
	for _, r := range s {
		for _, p := range pairs {
			temp, _ := stack.Peek()
			if r == p.open {
				stack.Push(r)
				break
			} else if r == p.close && stack.IsEmpty() {
				return false
			} else if r == p.close && temp == p.open {
				stack.Pop()
				break
			} else if r == p.close && temp != p.open {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func (stack *Stack) reverseStack() {
	if stack.IsEmpty() {
		return
	}

	data, _ := stack.Pop()
	stack.reverseStack()
	stack.insertAtBottom(data)
}

func (stack *Stack) insertAtBottom(data interface{}) {
	if stack.IsEmpty() {
		stack.Push(data)
		return
	}

	temp, _ := stack.Pop()
	stack.insertAtBottom(data)
	stack.Push(temp)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pairWiseConsecutive(stack *Stack) bool {
	auxStack := NewStack(1)
	for !stack.IsEmpty() {
		peek, _ := stack.Peek()
		auxStack.Push(peek)
		stack.Pop()
	}

	result := true
	for auxStack.Size() > 1 {
		x, _ := auxStack.Peek()
		auxStack.Pop()
		y, _ := auxStack.Peek()
		auxStack.Pop()

		if abs(x.(int)-y.(int)) != 1 {
			result = false
		}

		stack.Push(x)
		stack.Push(y)
	}

	if auxStack.Size() == 1 {
		peek, _ := auxStack.Peek()
		stack.Push(peek)
	}
	return result
}

func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			// remove
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
