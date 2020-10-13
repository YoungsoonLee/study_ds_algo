package main

import (
	"fmt"

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

func main() {
	input := []int{3, 4, 7, 6, 5, 9, 10}

	fmt.Println("bubble: ", bubble.BubbleSort2(input))
	fmt.Println("selection: ", selection.SelectionSort(input))
	fmt.Println("insertion: ", insertion.InsertionSort(input))
	fmt.Println("shell: ", shell.ShellSort(input))
	fmt.Println("merge: ", merge.MergeSort(input))
	fmt.Println("quick: ", quick.QuickSort(input))
}
