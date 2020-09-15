package ddlinkedList

type DLLNode struct {
	Val  int
	Prev *DLLNode
	Next *DLLNode
}

type DLL struct {
	head *DLLNode
	tail *DLLNode
	size int
}

func (dll *DLL) CheckIfEmptyAndAdd(newNode *DLLNode) bool {
	if dll.size == 0 {
		dll.head = newNode
		dll.tail = newNode
		dll.size++
		return true
	}

	return false
}

// TC: O(1), SC: O(1)
func (dll *DLL) InsertBeginning(val int) {
	newNode := &DLLNode{Val: val, Prev: nil, Next: nil}
	if !dll.CheckIfEmptyAndAdd(newNode) {
		head := dll.head
		newNode.Next = head
		newNode.Prev = nil

		head.Prev = newNode

		dll.head = newNode
		dll.size++
		return
	}
	return
}

func (dll *DLL) InsertEnd(val int) {
	newNode := &DLLNode{Val: val, Prev: nil, Next: nil}

	if !dll.CheckIfEmptyAndAdd(newNode) {
		head := dll.head
		for i := 0; i < dll.size; i++ {
			if head.Next == nil {
				newNode.Prev = head
				newNode.Next = nil

				head.Next = newNode

				dll.tail = newNode
				dll.size++
				break
			}

			head = head.Next

		}
	}

	return
}

func (dll *DLL) Insert(val int, loc int) {
	newNode := &DLLNode{Val: val, Prev: nil, Next: nil}

	if !dll.CheckIfEmptyAndAdd(newNode) {
		head := dll.head
		for i := 1; i < dll.size; i++ {
			if i == loc {
				newNode.Prev = head.Prev
				newNode.Next = head

				head.Prev.Next = newNode
				head.Prev = newNode

				dll.size++
				return
			}

			head = head.Next
		}
	}
	return
}

func (dll *DLL) DeleteFirst() int {
	// check empty

	head := dll.head
	if head.Prev == nil {
		deletedNode := head.Val

		dll.head = head.Next
		dll.head.Prev = nil

		return deletedNode
	}

	return -1
}

func (dll *DLL) DeleteLast() int {
	// check empty

	head := dll.head

	for {
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	dll.tail = head.Prev
	dll.tail.Next = nil
	dll.size--
	return head.Val
}

func (dll *DLL) Delete(pos int) int {
	// check empty

	head := dll.head
	for i := 1; i <= pos; i++ {
		if head.Next == nil && pos > i {
			// list is lesser then given position
			return -1
		} else if i == pos {
			head.Prev.Next = head.Next
			head.Next.Prev = head.Prev
			dll.size--
			return head.Val
		}

		head = head.Next
	}
	return -1
}
