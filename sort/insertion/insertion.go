package insertion

func InsertionSort(A []int) []int {
	n := len(A)

	for i := 1; i <= n-1; i++ {
		j := i
		for j > 0 {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
			j -= 1
		}
	}
	return A
}
