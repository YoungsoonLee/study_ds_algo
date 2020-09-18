package models

import (
	"errors"
	"math"
)

type MultiStackes struct {
	top1, top2 int
	capacity   int
	array      []int
}

func (stacks *MultiStackes) Init(capacity int) *MultiStackes {
	stacks.top1 = -1
	stacks.top2 = capacity
	stacks.capacity = capacity
	stacks.array = make([]int, capacity)
	return stacks
}

func NewStackMultiStackes(capacity int) *MultiStackes {
	return new(MultiStackes).Init(capacity)
}

func (stacks *MultiStackes) Size(stakNumber int) int {
	if stakNumber == 1 {
		return stacks.top1 + 1
	} else {
		return stacks.capacity - stacks.top2
	}
}

func (stacks *MultiStackes) IsFull() bool {
	return (stacks.Size(1) + stacks.Size(1)) == stacks.capacity
}

func (stacks *MultiStackes) IsEmpty(stackNumber int) bool {
	if stackNumber == 1 {
		return stacks.top1 == -1
	} else {
		return stacks.top2 == stacks.capacity
	}
}

func (stacks *MultiStackes) Push(stackNumber int, data int) error {
	if stacks.IsFull() {
		return errors.New("stackes is full")
	}

	if stackNumber == 1 {
		stacks.top1++
		stacks.array[stacks.top1] = data
	} else {
		stacks.top2 = stacks.top2 - 1
		stacks.array[stacks.top2] = data
	}
	return nil
}

func (stacks *MultiStackes) Pop(stackNumber int) int {
	var result int
	if stacks.IsEmpty(stackNumber) {
		return math.MinInt32
	}
	if stackNumber == 1 {
		result = stacks.array[stacks.top1]
		stacks.top1--
	} else {
		result = stacks.array[stacks.top2]
		stacks.top2++
	}
	return result
}
