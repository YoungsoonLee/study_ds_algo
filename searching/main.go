package main

import (
	"fmt"
	"sort"
)

func binarySearch(data int, A []int) bool {
	low, high := 0, len(A)-1
	for low <= high {
		mid := (low + high) / 2
		if A[mid] < data {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(A) || A[low] != data {
		return false
	}
	return true
}

func binarySearchRecursive(data int, A []int) bool {
	low, high := 0, len(A)-1
	if low <= high {
		mid := (low + high) / 2
		if A[mid] > data {
			return binarySearchRecursive(data, A[:mid])
		} else if A[mid] < data {
			return binarySearchRecursive(data, A[mid+1:])
		} else {
			return true
		}
	}
	return false
}

func interpolationSearch(A []int, data int) int {
	low, high := 0, len(A)-1
	for low <= high && data >= A[low] && data <= A[high] {
		guessIndex := low + (data-A[low])*(high-low)/(A[high]-A[low])
		if A[guessIndex] == data {
			return guessIndex
		} else if A[guessIndex] < data {
			low = guessIndex + 1
		} else if A[guessIndex] > data {
			high = guessIndex - 1
		}
	}

	return -1
}

func checkDuplicateInArray(A []int) bool {
	var HT = map[int]bool{}
	for _, num := range A {
		if _, ok := HT[num]; ok {
			return true
		}
		HT[num] = true
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// best
/*
Notes:
 This solution does not work if the given array is read only.
 This solution will work only if all the array elements are positive.
 If the elements range is not in 0 to 􏰂 − 1 then it may give exceptions.
*/
func checkDuplicateInArray2(A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[abs(A[i])] < 0 {
			return true
		} else {
			A[abs(A[i])] = -A[abs(A[i])]
		}
	}
	return false
}

func maxRepetitions(A []int) []int {
	maxCounter, maxElement, n := 0, -1, len(A)
	for i := 0; i < n; i++ {
		A[A[i]%n] += n
	}

	for i := 0; i < n; i++ {
		if A[i]/n > maxCounter {
			maxCounter = A[i] / n
			maxElement = i
		}
	}

	return []int{maxCounter, maxElement}
}

func missingNumber(A []int) int {
	for v := 0; v <= len(A); v++ {
		found := false
		for i := 0; i < len(A); i++ {
			if v == A[i] {
				found = true
				break
			}
		}
		if found == false {
			return v
		}
	}
	return -1
}

func missingNumber2(A []int) int {
	if len(A) == 0 {
		return -1
	}

	sort.Ints(A)

	if A[len(A)-1] != len(A) {
		return len(A)
	}

	for i := 0; i < len(A); i++ {
		if A[i] != i {
			return i
		}
	}
	return -1
}

func missingNumber3(A []int) int {
	var HT = map[int]bool{}
	for _, num := range A {
		HT[num] = true
	}

	for v := 0; v <= len(A); v++ {
		if _, ok := HT[v]; ok == false {
			return v
		}
	}
	return -1
}

func missingNumberXOR(A []int) int {
	X, Y := 0, 0
	for i := 0; i < len(A); i++ {
		X ^= A[i]
	}

	for i := 0; i < len(A); i++ {
		Y ^= i
	}

	return X ^ Y
}

func singleNumber4(A []int) int {
	result := 0
	for _, n := range A {
		result ^= n
		fmt.Println("result: ", result)
	}
	return result
}

func repeatedElementsHashing(A []int) {
	counts := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		counts[A[i]]++
		if counts[A[i]] == 2 {
			fmt.Println(A[i])
		}
	}
}

func twoSum(A []int, target int) []int {
	for i, v := range A {
		for j := i + 1; j < len(A); j++ {
			if A[j] == target-v {
				return []int{i, j}
			}
		}
	}
	panic("should never happen")
}

func twoSum2(A []int, target int) []int {
	sort.Ints(A)
	fmt.Println(A)
	i, j := 0, len(A)-1
	for i < j {
		sum := A[i] + A[j]
		if sum == target {
			break
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return []int{i, j}
}

func twoSum3(A []int, K int) []int {
	H := make(map[int]int)
	for i, v := range A {
		k := K - v
		if _, ok := H[k]; ok {
			return []int{H[k], i}
		}
		H[v] = i
	}
	return nil
}

func twoElementsWithMinSumCloseToZero(A []int) []int {
	n := len(A)
	if n < 2 {
		return []int{}
	}

	min_i, min_j := 0, 1
	minSum := A[0] + A[1]
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := A[i] + A[j]
			if abs(minSum) > abs(sum) {
				minSum = sum
				min_i = i
				min_j = j
			}
		}
	}

	return []int{A[min_i], A[min_j]}

}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(items, 63))

	items = []int{1, 2, 3, 2, 3, 1, 3}
	fmt.Println(singleNumber4(items))
}
