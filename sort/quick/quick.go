package quick

func QuickSort(A []int) []int {
	recursionSort(A, 0, len(A)-1)
	return A
}

func recursionSort(A []int, left, right int) {
	if left < right {
		pivot := partition(A, left, right)
		recursionSort(A, left, pivot-1)
		recursionSort(A, pivot+1, right)
	}
}

func partition(A []int, left int, right int) int {
	for left < right {
		for left < right && A[left] <= A[right] {
			right--
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			left++
		}
		for left < right && A[left] <= A[right] {
			left++
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			right--
		}
	}
	return left
}
