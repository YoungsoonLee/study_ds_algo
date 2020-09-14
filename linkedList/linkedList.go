package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

func (ll *LinkedList) Display() error {
	if ll.Head == nil {
		return fmt.Errorf("List is empty")
	}

	current := ll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Val)
		current = current.Next
	}

	fmt.Println()
	return nil
}

func (ll LinkedList) Length() int {
	return ll.Size
}

func (ll *LinkedList) InsertBeginning(val int) {
	node := &ListNode{Val: val}

	if ll.Head == nil {
		ll.Head = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}
	ll.Size++
	return
}

func (ll *LinkedList) InsertEnd(val int) {
	node := &ListNode{Val: val}
	if ll.Head == nil {
		ll.Head = node
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}

		current.Next = node
	}

	ll.Size++
	return
}

func (ll *LinkedList) InsertPosition(position int, val int) error {
	if position < 1 || position > ll.Size+1 {
		return fmt.Errorf("Insert: out of postion")
	}
	node := &ListNode{Val: val}

	var prev, curr *ListNode
	prev = nil
	curr = ll.Head

	for position > 1 {
		prev = curr
		curr = curr.Next
		position = position - 1
	}

	if prev != nil {
		prev.Next = node
		node.Next = curr
	}

	ll.Size++
	return nil
}

func (ll *LinkedList) DeleteFirst() (int, error) {
	if ll.Head == nil {
		return 0, fmt.Errorf("List is empty")
	}

	data := ll.Head.Val
	ll.Head = ll.Head.Next
	ll.Size--

	return data, nil
}

func (ll *LinkedList) DeleteLast() (int, error) {
	if ll.Head == nil {
		return 0, fmt.Errorf("List is emoty")
	}

	var prev *ListNode
	curr := ll.Head
	for curr != nil {
		prev = curr
		curr = curr.Next
	}

	if prev != nil {
		prev.Next = nil
	} else {
		ll.Head = nil
	}

	ll.Size--
	return curr.Val, nil

}

func (ll *LinkedList) Delete(position int) (int, error) {
	if position < 1 || position > ll.Size {
		return 0, fmt.Errorf("out of index")
	}

	var prev, curr *ListNode
	prev = nil
	curr = ll.Head
	pos := 0

	if position == 1 {
		ll.Head = ll.Head.Next
	} else {
		for pos != position-1 {
			pos = pos + 1
			prev = curr
			curr = curr.Next
		}

		if curr != nil {
			prev.Next = curr.Next
		}
	}

	ll.Size--
	return curr.Val, nil
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

func findLoopBeginning(head *ListNode) *ListNode {
	fast, slow := head, head
	loopExists := false

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			loopExists = true
			break
		}
	}

	if loopExists {
		slow = head
		for slow != fast {
			fast = fast.Next
			slow = slow.Next
		}
		return slow
	}
	return nil
}

func findLoopLength(head *ListNode) int {
	fast, slow := head, head
	loopExists := false

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			loopExists = true
			break
		}
	}

	if loopExists {
		counter := 1
		fast = fast.Next
		for slow != fast {
			fast = fast.Next
			counter++
		}
		return counter
	}
	return 0
}

func (ll *LinkedList) sortedInsert(val int) {
	newNode := &ListNode{Val: val}

	// special case for the head end
	if ll.Head == nil || ll.Head.Val >= val {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil && current.Next.Val < val {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
}

// reverse !!!
func reverseList(head *ListNode) *ListNode {
	var prev, current *ListNode
	for current = head; current != nil; {
		current.Next, prev, current = prev, current, current.Next
	}
	return prev
}

func reverList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	h := reverse(head)
	head.Next = nil
	return h
}

func reverse(current *ListNode) *ListNode {
	if current == nil {
		return nil
	}

	temp := reverse(current.Next)
	if temp == nil {
		return current
	}

	current.Next.Next = current
	return temp
}

// brute force
// O(mnn)
func getInstersactionNode(head1, head2 *ListNode) *ListNode {
	for head1 != nil {
		temp := head2
		for temp != nil {
			if temp == head1 {
				return head1
			}
			temp = temp.Next
		}
		head1 = head1.Next
	}
	return nil
}

func getInstersactionNode2(head1, head2 *ListNode) *ListNode {
	len1, len2 := findLen(head1), findLen(head2)
	if len1 > len2 {
		for ; len1 > len2; len1-- {
			head1 = head1.Next
		}
	} else {
		for ; len2 > len1; len2-- {
			head2 = head2.Next
		}
	}

	for head1 != head2 {
		head1, head2 = head1.Next, head2.Next
	}
	return head1
}

func findLen(head *ListNode) int {
	l := 0
	for ; head != nil; head = head.Next {
		l++
	}
	return l
}
