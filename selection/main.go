package main

import (
	"container/heap"
	"fmt"
	"math"
)

func findMax(A []int) (max int) {
	max = math.MinInt32
	for _, value := range A {
		if value > max {
			max = value
		}
	}
	return max
}

func findMinMaxWithPairComparison(A []int) (min, max int) {
	min, max, n := math.MaxInt32, math.MinInt32, len(A)
	var i int
	for i := 0; i < n-1; i += 2 {
		if A[i] < A[i+1] {
			if A[i] < min {
				min = A[i]
			}
			if A[i+1] > max {
				max = A[i+1]
			}
		} else {
			if A[i+1] < min {
				min = A[i]
			}
			if A[i] > max {
				max = A[i]
			}
		}
	}

	if n%2 == 1 {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}
	return min, max
}

func findLargestAndSeondLargest(A []int) (largest, second int) {
	largest, second = math.MinInt32, math.MinInt32
	for _, v := range A {
		if v > largest {
			second = largest
			largest = v
		} else if v > second {
			second = v
		}
	}
	return largest, second
}

type Heap struct {
	items []int
}

func (h *Heap) Len() int {
	return len(h.items)
}

func (h *Heap) Less(i, j int) bool {
	if h.items[i] < h.items[j] {
		return true
	}
	return false
}

func (h *Heap) Pop() interface{} {
	old := h
	n := len(old.items)
	x := old.items[n-1]
	h.items = old.items[0 : n-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	h.items = append(h.items, x.(int))
}

func (h *Heap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func findKSmallest(S []int, k int) []int {
	h := &Heap{items: []int{}}
	heap.Init(h)
	result := []int{}
	for i := 0; i < len(S); i++ {
		heap.Push(h, S[i])
	}

	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(int))
	}
	return result
}

func findKSmallest2(S []int, k int) []int {
	for start, end := 0, len(S); ; {
		idx := partition(S, start, end)
		if idx == k-1 {
			break
		} else if idx > k-1 {
			end = idx
		} else {
			start = idx + 1
		}
	}

	return S[:k]
}

func partition(S []int, start, end int) int {
	pivot := start
	S[end-1], S[pivot] = S[pivot], S[end-1]
	idx := start
	for i := start; i < end-1; i++ {
		if S[i] < S[end-1] {
			S[i], S[idx] = S[idx], S[i]
			idx++
		}
	}
	S[end-1], S[idx] = S[idx], S[end-1]
	return idx
}

func findMedianSortedArrays(A, B []int) float32 {
	m, n := len(A), len(B)
	if m > n {
		A, B = B, A
		m, n = n, m
	}

	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && B[j-1] > A[i] {
			iMin = i + 1 // too small
		} else if i > iMin && A[i-1] > B[j] {
			iMax = i - 1 // too big
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = B[j-1]
			} else if j == 0 {
				maxLeft = A[i-1]
			} else {
				maxLeft = max(A[i-1], B[j-1])
			}

			if (m+n)%2 == 1 {
				return float32(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = B[j]
			} else if j == n {
				minRight = A[i]
			} else {
				minRight = min(B[j], A[i])
			}
			return float32((maxLeft + minRight) / 2.0)
		}
	}
	return 0.0
}

func main() {
	fmt.Println(findKSmallest([]int{11, -4, 7, 8, -10}, 3))
}
