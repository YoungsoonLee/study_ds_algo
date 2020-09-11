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
