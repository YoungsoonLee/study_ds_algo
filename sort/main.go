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

func main() {
	input := []int{3, 4, 7, 6, 5, 9, 10}

	fmt.Println("bubble: ", bubble.BubbleSort2(input))
	fmt.Println("selection: ", selection.SelectionSort(input))
	fmt.Println("insertion: ", insertion.InsertionSort(input))
	fmt.Println("shell: ", shell.ShellSort(input))
	fmt.Println("merge: ", merge.MergeSort(input))
	fmt.Println("quick: ", quick.QuickSort(input))
}
