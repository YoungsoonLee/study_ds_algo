package linkedList

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

// O(n) + O(n) = O(n)
func middleNode(head *ListNode) *ListNode {
	l := findLen(head)
	count, target := 0, (l/2)+1

	for {
		count++
		if count == target {
			return head
		}
		head = head.Next
	}
}

func middleNode2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// O(n),  recursive is stack !!! O(n)
func printListInReverse(head *ListNode) {
	if head == nil {
		return
	}

	printListInReverse(head.Next)
	fmt.Print(head.Val)
}

// O(n/2) = O(n)
func (ll *LinkedList) IsLengthEven() bool {
	current := ll.Head
	for current != nil && current.Next != nil {
		current = current.Next.Next
	}

	if current != nil {
		return false
	}
	return true
}

// O(n+m)
// recursive
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// iterative
func mergeTwoLists_iter(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)

	for node := dummy; l1 != nil || l2 != nil; node = node.Next {
		if l1 == nil {
			node.Next = l2
			break
		} else if l2 == nil {
			node.Next = l1
			break
		} else if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

	}
	return dummy.Next
}

func mergeKLists(list []*ListNode) *ListNode {
	if list == nil || len(list) == 0 {
		return nil
	}

	for len(list) > 1 {
		l1 := list[0]
		l2 := list[1]
		list = list[2:]

		merged := mergeTwoLists(l1, l2)
		list = append(list, merged)
	}
	return list[0]
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	firstTail := slow
	slow = slow.Next
	firstTail.Next = nil // divide the first list and the second

	first, second := sortList(head), sortList(slow)
	return merge(first, second)
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	curHead := &ListNode{}
	tmpHead := curHead

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			curHead.Next = head1
			head1 = head1.Next
			curHead = curHead.Next
		} else {
			curHead.Next = head2
			head2 = head2.Next
			curHead = curHead.Next
		}
	}

	// remains
	if head1 != nil {
		curHead.Next = head1
	} else if head2 != nil {
		curHead.Next = head2
	}

	return tmpHead.Next
}

func splitList(head *ListNode) (head1 *ListNode, head2 *ListNode) {
	var slow, fast *ListNode
	if head == nil || head.Next == nil { // length < 2
		head1 = head
		head2 = nil
	} else {
		slow = head
		fast = head.Next

		for fast != nil {
			fast = fast.Next
			if fast != nil {
				slow = slow.Next
				fast = fast.Next
			}
		}
		//
		head1 = head
		head2 = slow.Next
		slow.Next = nil
	}
	return head1, head2
}

// hmm...
func reverseBlockOfKNodes(head *ListNode, k int) *ListNode {
	// base check
	if head == nil || k == 1 {
		return head
	}

	// get length
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}

	// result
	result := ListNode{0, head}
	prev := &result

	for step := 0; step+k <= length; step = step + k {
		tail := prev.Next
		nextNode := tail.Next

		for i := 1; i < k; i++ {
			tail.Next = nextNode.Next
			nextNode.Next = prev.Next
			prev.Next = nextNode
			nextNode = tail.Next
		}

		prev = tail
	}

	return result.Next
}

// iterative reversePairs
// !!!!!
func reversePairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := head.Next
	var prev *ListNode
	for head != nil && head.Next != nil {
		nextNode := head.Next
		head.Next = nextNode.Next
		nextNode.Next = head
		if prev != nil {
			prev.Next = nextNode
		}
		prev = head
		head = head.Next
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, result := 0, new(ListNode)
	for node := result; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{carry % 10, nil}
		carry /= 10
	}
	return result.Next
}
