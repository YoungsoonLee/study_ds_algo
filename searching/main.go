package main

import (
	"fmt"
	"math"
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

func threeSum(A []int, target int) [][]int {
	var result [][]int
	sort.Ints(A)
	for i := 0; i < len(A)-2; i++ {
		if i > 0 && A[i] == A[i-1] {
			continue
		}

		target, left, right := target-A[i], i+1, len(A)-1
		for left < right {
			sum := A[left] + A[right]
			if sum == target {
				result = append(result, []int{A[i], A[left], A[right]})
				left++
				right--
				for left < right && A[left] == A[left-1] {
					left++
				}
				for left < right && A[right] == A[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}

		}
	}
	return result
}

func search(A []int, data int) int {
	left := 0
	right := len(A) - 1

	for left < right {
		mid := left + (right-left)/2
		if A[mid] == data {
			return mid
		}

		if A[mid] >= A[left] {
			if data <= A[mid] && data >= A[left] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if data > A[mid] && data < A[left] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	if left >= 0 && left <= len(A)-1 && A[left] == data {
		return left
	}
	if right >= 0 && right <= len(A)-1 && A[right] == data {
		return right
	}
	return -1
}

func isMonotonic(A []int) bool {
	if len(A) < 2 {
		return true
	}

	isIncreasing := 0

	for i := 1; i < len(A); i++ {
		if isIncreasing == 0 {
			if A[i-1] > A[i] {
				isIncreasing = -1
			} else if A[i-1] < A[i] {
				isIncreasing = 1
			}
		}
		if isIncreasing == 1 && A[i-1] > A[i] {
			return false
		}
		if isIncreasing == -1 && A[i-1] < A[i] {
			return false
		}
	}
	return true
}

func separateEvenOdd(A []int) {
	left, right := 0, len(A)-1
	for left < right {
		for A[left]%2 == 0 && left < right {
			left++
		}
		for A[right]%2 == 1 && left < right {
			right--
		}

		if left < right {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		}
	}

	fmt.Println(A)
}

// it's only 0 or 1
func separate0and1(A []int) []int {
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			count++
		}
	}

	for i := 0; i < count; i++ {
		A[i] = 0
	}

	for i := count; i < len(A); i++ {
		A[i] = 1
	}
	return A
}

func separate0and1_2(A []int) []int {
	left, right := 0, len(A)-1
	for left < right {
		for A[left] == 0 && left < right {
			left++
		}

		for A[right] == 1 && left < right {
			right--
		}

		if left < right {
			A[left] = 0
			A[right] = 1
			left++
			right--
		}
	}
	return A
}

func DutchFlagProblem(A []int) []int {
	low, mid, high := 0, 0, len(A)-1
	for mid <= high {
		switch A[mid] {
		case 0:
			A[low], A[mid] = A[mid], A[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			A[mid], A[high] = A[high], A[mid]
			high--
		}
	}
	return A
}

func shuffleArray1(A []int, left, right int) {
	n := len(A) / 2
	for i, q, k := 0, 1, n; i < n; i, k, q = i+1, k+1, q+1 {
		for j := k; j > i+q; j-- {
			A[j-1], A[j] = A[j], A[j-1]
		}
	}
}

func maxIndexDiff(A []int) int {
	maxDiff, n := -1, len(A)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if A[j] > A[i] && maxDiff < (j-i) {
				maxDiff = j - i
			}
		}
	}
	return maxDiff
}

func maxIndexDiff2(A []int) int {
	n := len(A)
	leftMins, rightMaxs := make([]int, n), make([]int, n)
	leftMins[0] = A[0]
	for i := 1; i < n; i++ {
		leftMins = min(A[i], leftMins[i-1])
	}

	rightMaxs[n-1] = A[n-1]
	for j := n - 2; j >= 0; j-- {
		rightMaxs[j] = max(A[j], rightMaxs[j+1])
	}

	i, j, maxDiff := 0, 0, -1
	for j < n && i < n {
		if leftMins[i] < rightMaxs[j] {
			maxDiff = max(maxDiff, j-i)
			j += 1
		} else {
			i += 1
		}
	}
	return maxDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkPairwiseSorted(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	for i := 0; i < len(A)-1; i += 2 {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true

}

func pivotIndex(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	leftSum := 0
	for i := 0; i < len(A); i++ {
		if leftSum == sum-A[i]-leftSum {
			return i
		}
		leftSum += A[i]
	}
	return -1
}

func replaceWithNearestGreaterElement(A []int) []int {
	n := len(A)
	for i := 0; i < n; i++ {
		nextNearestGreater := math.MinInt32
		for j := i + 1; j < n; j++ {
			if A[j] > nextNearestGreater {
				nextNearestGreater = A[j]
			}
		}
		A[i] = nextNearestGreater
	}
	return A
}

func replaceWithNearestGreaterElement2(A []int) []int {
	greatest := math.MinInt32
	for i := len(A) - 1; i >= 0; i-- {
		temp := greatest
		if A[i] > greatest {
			greatest = A[i]
		}
		A[i] = temp
	}

	return A
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(items, 63))

	items = []int{1, 2, 3, 2, 3, 1, 3}
	fmt.Println(singleNumber4(items))

	items = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	separateEvenOdd(items)

	items = []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	fmt.Println(DutchFlagProblem(items))
}
