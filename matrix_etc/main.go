package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O(n^2)
func spiral(A [][]int) {
	rowStart, rowEnd, columnStart, columnEnd := 0, len(A)-1, 0, len(A[0])-1
	for {
		// right
		for i := columnStart; i <= columnEnd; i++ {
			fmt.Printf("%d", A[rowStart][i])
		}
		rowStart++
		if rowStart > rowEnd {
			break
		}
		// down
		for i := rowStart; i <= rowEnd; i++ {
			fmt.Printf("%d", A[i][columnEnd])
		}
		columnEnd--
		if columnStart > columnEnd {
			break
		}
		// left
		for i := columnEnd; i >= columnStart; i-- {
			fmt.Printf("%d", A[rowEnd][i])
		}
		rowEnd--
		if rowStart > rowEnd {
			break
		}
		// up
		for i := rowEnd; i >= rowStart; i-- {
			fmt.Printf("%d", A[i][columnStart])
		}
		columnStart++
		if columnStart > columnEnd {
			break
		}
	}
}

// shuffle
func shuffle(cards []int, n int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

// rotate & reverse array
func leftRotate(arr []int, d int, n int) {
	if d == 0 {
		return
	}
	reverseArray(arr, 0, d-1)
	reverseArray(arr, d, n-1)
	reverseArray(arr, 0, n-1)
}

func printArray(arr []int, size int) {
	var i int
	for i = 0; i < size; i++ {
		fmt.Printf("%d", arr[i])
	}
}

func reverseArray(arr []int, start int, end int) {
	var temp int
	for start < end {
		temp = arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

func main() {
	A := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	spiral(A)
}
