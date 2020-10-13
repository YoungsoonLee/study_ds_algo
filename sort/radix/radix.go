package radix

func findLargestElement(A []int) int {
	largestElement := 0
	for i := 0; i < len(A); i++ {
		if A[i] > largestElement {
			largestElement = A[i]
		}
	}
	return largestElement
}

func RadixSort(A []int) {
	largestElement := findLargestElement(A)
	size := len(A)
	significantDigit := 1
	semiSorted := make([]int, size, size)
	for largestElement/significantDigit > 0 {
		bucket := [10]int{0}
		for i := 0; i < size; i++ {
			bucket[(A[i]/significantDigit)%10]++
		}
		for i := 0; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := size - 1; i >= 0; i-- {
			bucket[(A[i]/significantDigit)%10]--
			semiSorted[bucket[(A[i]/significantDigit)%10]] = A[i]
		}
		for i := 0; i < size; i++ {
			A[i] = semiSorted[i]
		}
		significantDigit *= 10
	}
}

func RadixSortTucker(A []int) {

}
