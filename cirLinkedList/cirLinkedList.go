package cirLinkedList

import "fmt"

type CLLNode struct {
	Val  int
	Next *CLLNode
}

type CLL struct {
	Head *CLLNode
	Size int
}

func (cll *CLL) Length() int {
	current := cll.Head
	count := 1

	if current == nil {
		return 0
	}

	current = current.Next
	for current != cll.Head {
		current = current.Next
		count++
	}
	return count
}

func (cll *CLL) Display() {
	head := cll.Head
	for i := 0; i < cll.Size; i++ {
		fmt.Print(head.Val)
		fmt.Print("-->")
		head = head.Next
	}

	fmt.Println()
}

func (cll *CLL) CheckIfEmptyAndAdd(val int) bool {
	newNode := &CLLNode{Val: val, Next: nil}

	if cll.Size == 0 {
		cll.Head = newNode
		cll.Head.Next = newNode // ???
		cll.Size++
		return true
	}
	return false
}

func (cll *CLL) InsertBeginning(val int) {
	if !cll.CheckIfEmptyAndAdd(val) {
		newNode := &CLLNode{Val: val, Next: nil}
		current := cll.Head
		newNode.Next = current

		for {
			if current.Next == cll.Head {
				break
			}
			current = current.Next
		}

		current.Next = newNode
		cll.Head = newNode // move haed !!!!!
		cll.Size++
	}

}

func (cll *CLL) InsertEnd(val int) {
	if !cll.CheckIfEmptyAndAdd(val) {
		newNode := &CLLNode{Val: val, Next: nil}

		current := cll.Head
		for {
			if current.Next == cll.Head {
				break
			}
			current = current.Next
		}

		current.Next = newNode
		newNode.Next = cll.Head
		cll.Size++
	}
}

func (cll *CLL) Insert(val int, pos int) {
	if !cll.CheckIfEmptyAndAdd(val) {
		current := cll.Head
		count := 1

		if pos == 1 {
			cll.InsertBeginning(val)
		}

		newNode := &CLLNode{Val: val, Next: nil}

		for {
			if current.Next == nil && pos-1 > count {
				break
			}
			if count == pos-1 {
				newNode.Next = current.Next
				current.Next = newNode
				cll.Size++
				break
			}
			current = current.Next
			count++
		}
	}
}

func (cll *CLL) CheckIfEmpty() bool {
	if cll.Size == 0 {
		return true
	}

	return false
}

func (cll *CLL) DeleteBeginning() int {
	if !cll.CheckIfEmpty() {
		current := cll.Head
		deletedElem := current.Val

		if cll.Size == 1 {
			cll.Head = nil
			cll.Size--
			return deletedElem
		}

		prevStart := cll.Head
		cll.Head = current.Next
		for {
			if current.Next == prevStart {
				break
			}
			current = current.Next
		}

		current.Next = cll.Head
		cll.Size--
		return deletedElem
	}
	return -1
}

func (cll *CLL) DeleteEnd() int {
	if !cll.CheckIfEmpty() {
		current := cll.Head
		deletedEle := current.Val

		if cll.Size == 1 {
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}

		for {
			if current.Next.Next == cll.Head {
				deletedEle = current.Next.Val
				break
			}
			current = current.Next
		}

		current.Next = cll.Head
		cll.Size--
		return deletedEle
	}
	return -1
}

func (cll *CLL) Delete(pos int) int {
	if !cll.CheckIfEmpty() {
		current := cll.Head
		deleteEle := current.Val
		if cll.Size == 1 {
			deleteEle = cll.DeleteBeginning()
			return deleteEle
		}

		if cll.Size == pos {
			deleteEle = cll.DeleteEnd()
			return deleteEle
		}

		count := 1
		for {
			if count == pos-1 {
				deleteEle = current.Next.Val
				break
			}
			current = current.Next
		}

		current.Next = current.Next.Next
		cll.Size--
		return deleteEle

	}
	return -1
}

// create a new node of circular linked list
func NewListNode(val int) *CLLNode {
	temp := &CLLNode{}
	temp.Next = temp
	temp.Val = val
	return temp
}

func getJosephusPosition(m, n int) {
	// create a circular linked list of size N
	head := NewListNode(1)
	prev := head

	for i := 2; i <= n; i++ {
		prev.Next = NewListNode(i)
		prev = prev.Next
	}

	prev.Next = head // connect last node to first

	ptr1, ptr2 := head, head
	for ptr1.Next != ptr1 {
		count := 1
		for count != m {
			ptr2 = ptr1
			ptr1 = ptr1.Next
			count++
		}

		// remomve the m-th node
		ptr2.Next = ptr1.Next
		ptr1 = ptr2.Next
	}
	fmt.Println("Last person left standing ", "(Josephus Position) is ", ptr1.Val)
}
