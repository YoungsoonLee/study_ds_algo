package shell

func ShellSort(A []int) []int {
	n := len(A)
	h := 1

	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && A[j] < A[j-h]; j -= h {
				A[j], A[j-h] = A[j-h], A[j]
			}
		}
		h /= 3
	}
	return A
}
