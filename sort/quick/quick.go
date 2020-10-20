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

func QuickSortTucker(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	p := len(arr) - 1
	i := 0
	j := p - 1

	for i <= j {
		if arr[i] < arr[p] {
			i++
		} else {
			if arr[j] >= arr[p] {
				j--
			} else {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
	}

	arr[p], arr[i] = arr[i], arr[p]
	QuickSortTucker(arr[:i])
	QuickSortTucker(arr[i+1:])
	return arr
}
