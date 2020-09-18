package models

type StackLL struct {
	top  *ListNode
	size int
}

type ListNode struct {
	data interface{}
	next *ListNode
}

func (s *StackLL) length() int {
	return s.size
}

func (s *StackLL) IsEmpty() bool {
	return s.size == 0
}

func (s *StackLL) Push(data interface{}) {
	s.top = &ListNode{data, s.top}
	s.size++
}

func (s *StackLL) Pop() (data interface{}) {
	if s.size > 0 {
		data, s.top = s.top.data, s.top.next
		s.size--
	}
	return nil
}

func (s *StackLL) Peek() (data interface{}) {
	if s.size > 0 {
		data = s.top.data
		return
	}
	return nil
}
