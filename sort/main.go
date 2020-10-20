package main

import (
	"fmt"
	"sort"

	"github.com/YoungsoonLee/study-ds-algo/sort/bubble"
	"github.com/YoungsoonLee/study-ds-algo/sort/insertion"
	"github.com/YoungsoonLee/study-ds-algo/sort/merge"
	"github.com/YoungsoonLee/study-ds-algo/sort/quick"
	"github.com/YoungsoonLee/study-ds-algo/sort/selection"
	"github.com/YoungsoonLee/study-ds-algo/sort/shell"
)

func CheckDuplicatesInArray(A []int) bool {
	for i := 0; i < len(A); i++ {
		for v := 0; v < i; v++ {
			if A[v] == A[i] {
				return true
				break
			}
		}
	}
	return false
}

func CheckDuplicatesInOrderedArray(A []int) bool {
	sort.Ints(A)
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return true
		}
	}
	return false
}

func CheckWhoWinsTheElection(A []int) int {
	maxCounter, counter, candidate := 0, 0, A[0]
	for i := 0; i < len(A); i++ {
		candidate = A[i]
		counter = 0

		for j := i + 1; j < len(A); j++ {
			if A[i] == A[j] {
				counter++
			}
		}

		if counter > maxCounter {
			maxCounter = counter
			candidate = A[i]
		}
	}
	return candidate
}

func CheckWhoWinsTheElectionWithSort(A []int) int {
	currentCounter, maxCounter, currentCandidate, maxCandidate := 1, 1, A[0], 0
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i] == currentCandidate {
			currentCounter++
		} else {
			currentCandidate = A[i]
			currentCounter = 1
		}

		if currentCounter > maxCounter {
			maxCandidate = currentCandidate
			maxCounter = currentCounter
		}
	}
	return maxCandidate
}

/*
func find(A, B []int, K int) bool {
	sort.Ints(A) // nlogn
	for i := 0; i < len(B); i++ {
		c := K - B[i]
		if binarySearch(c, A) {
			return true
		}
	}
	return false
}
*/

func mostFrequent(A []int) int {
	sort.Ints(A)
	currentCounter, maxCounter, res, n := 1, 1, A[0], len(A)
	for i := 1; i < n; i++ {
		if A[i] == A[i-1] {
			currentCounter++
		} else {
			if currentCounter > maxCounter {
				maxCounter = currentCounter
				res = A[i-1]
			}
			currentCounter = 1
		}
	}
	if currentCounter > maxCounter { // If last element is most frequent
		maxCounter = currentCounter
		res = A[n-1]
	}
	return res
}

func convertArraySawToothWave(A []int) {
	n := len(A)
	sort.Ints(A)
	for i := 1; i < n; i += 2 {
		if i+1 < n {
			A[i], A[i+1] = A[i+1], A[i]
		}
	}
}

func main() {

	input := []int{3, 4, 7, 6, 5, 9, 10}

	fmt.Println("bubble: ", bubble.BubbleSort2(input))
	fmt.Println("selection: ", selection.SelectionSort(input))
	fmt.Println("insertion: ", insertion.InsertionSort(input))
	fmt.Println("shell: ", shell.ShellSort(input))
	fmt.Println("merge: ", merge.MergeSort(input))
	fmt.Println("quick: ", quick.QuickSort(input))

	fmt.Println("quick: ", quick.QuickSortTucker(input))
}
