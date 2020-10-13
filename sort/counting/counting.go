package counting

import "fmt"

func CountingSort(A []int, K int) []int {
	bucketLen := K + 1

	C := make([]int, bucketLen)

	sortedIndex := 0
	length := len(A)

	for i := 0; i < length; i++ {
		C[A[i]] += 1
	}

	fmt.Println(C)

	for j := 0; j < bucketLen; j++ {
		for C[j] > 0 {
			A[sortedIndex] = j
			sortedIndex += 1
			C[j] -= 1
		}
	}
	return A
}
