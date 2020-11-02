package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Item interface {
	Less(item Item) bool
}

type Heap struct {
	size int
	data []Item
}

func New() *Heap {
	return &Heap{}
}

func parent(i int) int {
	return int(math.Floor(float64(i-1) / 2.0))
}

func leftChild(parent int) int {
	return (2 * parent) + 1
}

func rightChild(parent int) int {
	return (2 * parent) + 2
}

func getMinimum(h *Heap) (Item, error) {
	if h.size == 0 {
		return nil, fmt.Errorf("Unable to get element from empty heap")
	}

	return h.data[0], nil
}

func swap(h *Heap, i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// for insert
func (h *Heap) percolateUp() {
	idx := h.size - 1
	if idx <= 0 {
		return
	}

	for {
		p := parent(idx)
		if p < 0 || h.data[p].Less(h.data[idx]) {
			break
		}
		swap(h, p, idx) // swap
		idx = p         // !!!
	}
}

// for deleting
func (h *Heap) percolateDown(i int) {
	p := i
	for {
		l := leftChild(p)
		r := rightChild(p)
		s := p
		if l < h.size && h.data[l].Less(h.data[s]) {
			s = l
		}
		if r < h.size && h.data[r].Less(h.data[s]) {
			s = r
		}
		if s == p {
			break
		}
		swap(h, p, s)
		p = s
	}
}

// deleting
func (h *Heap) Extract() (Item, error) {
	n := h.size
	if n == 0 {
		return nil, fmt.Errorf("Unable to extract from emmpty Heap")
	}

	m := h.data[0]
	h.data[0] = h.data[n-1] // 끝에꺼 올림
	h.data = h.data[:n-1]
	h.size--
	if h.size > 0 {
		h.percolateDown(0) // 다시 heap 만듬
	} else {
		h.data = nil
	}
	return m, nil
}

// insert
func (h *Heap) Insert(item Item) {
	if h.size == 0 {
		h.data = make([]Item, 1)
		h.data[0] = item
	} else {
		h.data = append(h.data, item)
	}
	h.size++
	h.percolateUp() //
}

func Heapify(items []Item) *Heap {
	h := New()
	n := len(items)
	h.data = make([]Item, n)
	copy(h.data, items)
	h.size = len(items)
	i := int(n / 2)
	for i >= 0 {
		h.percolateDown(i)
		i--
	}
	return h
}

// below code if for testing
type Int int

func (a Int) Less(b Item) bool {
	val, ok := b.(Int)
	return ok && a < val
}

func verifyHeap(h *Heap) bool {
	queue := make([]Int, 1)
	queue[0] = 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		l := leftChild(int(p))
		r := rightChild(int(p))
		if l < h.size {
			if !h.data[p].Less(h.data[l]) {
				return false
			}
			queue = append(queue, Int(l))
		}
		if r < h.size {
			if !h.data[p].Less(h.data[r]) {
				return false
			}
			queue = append(queue, Int(r))
		}
	}
	return true
}

func verifyStriclyIncreasing(h *Heap) (bool, []Item) {
	prev, _ := h.Extract()
	order := []Item{prev}
	for h.size > 0 {
		curr, _ := h.Extract()
		order = append(order, curr)
		if curr.Less(prev) {
			return false, order
		}
		prev = curr
		order = append(order, prev)
	}
	return true, order
}

func randomPerm(n int) []Item {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	ints := r.Perm(n)
	items := make([]Item, n)
	for idx, item := range ints {
		items[idx] = Int(item)
	}
	return items
}

func HeapSort(data []Item) []Item {
	hp := Heapify(data)
	size := len(hp.data)
	for i := size - 1; i > 0; i-- {
		swap(hp, 0, i)
		hp.size--
		hp.percolateDown(0)
	}
	hp.size = size
	return hp.data
}

func main() {
	items := randomPerm(30)
	for i := 0; i < len(items); i++ {
		fmt.Print(items[i].(Int), " ")
	}

	items = HeapSort(items)

	fmt.Print("\n")
	for i := 0; i < len(items); i++ {
		fmt.Print(items[i].(Int), " ")
	}

	/*
		items := randomPerm(10)
		hp := New()
		for _, item := range items {
			fmt.Println("Inserting an element into Heap: ", hp.data)
			hp.Insert(item)
		}

		if !verifyHeap(hp) {
			fmt.Println("invalid Heap: ", hp.data)
			return
		}

		if ok, order := verifyStriclyIncreasing(hp); !ok {
			fmt.Println("invalid Heap extraction order: ", order)
		}
	*/
}
