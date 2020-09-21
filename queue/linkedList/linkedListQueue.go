package linkedList

import (
	"bytes"
	"errors"
	"fmt"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

type Queue struct {
	front *ListNode
	rear  *ListNode
	size  int
}

func (q *Queue) isEmpty() bool {
	return q.front == nil
}

func (q *Queue) enQueue(data interface{}) {
	rear := new(ListNode)
	rear.data = data

	if q.isEmpty() {
		q.front = rear
	} else {
		oldLast := q.rear
		oldLast.next = rear
	}
	q.rear = rear
	q.size++
}

func (q *Queue) deQueue() (interface{}, error) {
	if q.isEmpty() {
		q.rear = nil
		return nil, errors.New("Unable to dequeue elemment. queue is empty")
	}
	data := q.front.data
	q.front = q.front.next
	q.size--
	return data, nil
}

func (q *Queue) frontElement() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("unable to peek element. queue is empty")
	}
	return q.front.data, nil
}

func (q *Queue) length() int {
	return q.size
}

func (q *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.size; i++ {
		result.WriteString(fmt.Sprintf("%v", j.data))
		if i < q.size-1 {
			result.WriteByte(' ')
		}
		j = j.next // !!!
	}
	result.WriteByte(']')
	return result.String()
}
