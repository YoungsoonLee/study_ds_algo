package stackLinkedList

type Stack struct {
	top  *ListNode
	size int
}

type ListNode struct {
	data interface{}
	next *ListNode
}

func (s *Stack) length() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(data interface{}) {
	s.top = &ListNode{data, s.top}
	s.size++
}

func (s *Stack) Pop() (data interface{}) {
	if s.size > 0 {
		data, s.top = s.top.data, s.top.next
		s.size--
	}
	return nil
}

func (s *Stack) Peek() (data interface{}) {
	if s.size > 0 {
		data = s.top.data
		return
	}
	return nil
}
