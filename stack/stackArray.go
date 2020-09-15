package stackArray

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
