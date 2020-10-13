package bubble

func BubbleSort(A []int) []int {
	n := len(A)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n-1; j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
	return A
}

func BubbleSort2(A []int) []int {
	var sorted bool
	items := len(A)
	for !sorted {
		sorted = true
		for i := 1; i < items; i++ {
			if A[i-1] > A[i] {
				A[i-1], A[i] = A[i], A[i-1]
				sorted = false
			}
		}
	}
	return A
}
