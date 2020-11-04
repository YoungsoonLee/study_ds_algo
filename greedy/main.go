package main

import (
	"math"
	"sort"
)

func maxEvents(start, fininsh []int) int {
	sort.Slice(start, func(i, j int) bool {
		if fininsh[i] == fininsh[j] {
			return start[i] < start[j]
		}
		return fininsh[i] < fininsh[j]
	})

	m := make(map[int]bool)
	for j := 0; j < len(start); j++ {
		for i := start[j]; i <= fininsh[j]; i++ {
			if _, ok := m[i]; !ok {
				m[i] = true
				break
			}
		}
	}
	return len(m)
}

func stockStrategy(A []int) []int {
	buyDateIndex, sellDateIndex, profit, n := 0, 0, 0, len(A)
	for i := 0; i < n; i++ { // buy
		for j := i + 1; j < n; j++ { //sell
			if A[j]-A[i] > profit {
				profit = A[j] - A[i]
				buyDateIndex = i
				sellDateIndex = j
			}
		}
	}
	return []int{profit, buyDateIndex, sellDateIndex}
}

func stockStrategy2(price []int, left, right int) (int, int, int) {
	if left == right {
		return 0, left, right
	}

	mid := left + (right-left)/2

	leftProfit, leftBuyDateIndex, leftSellDateIndex := stockStrategy2(price, left, right)
	rightProfit, rightBuyDateIndex, rightSellDateIndex := stockStrategy2(price, mid+1, right)

	minIndexLeft := minIndex(price, left, mid)
	maxIndexRight := maxIndex(price, mid+1, right)

	centerProfit := price[maxIndexRight] - price[minIndexLeft]
	if (centerProfit > leftProfit) && (centerProfit > rightProfit) {
		return centerProfit, minIndexLeft, maxIndexRight
	} else if (leftProfit > centerProfit) && (leftProfit > rightProfit) {
		return leftProfit, leftBuyDateIndex, leftSellDateIndex
	} else {
		return rightProfit, rightBuyDateIndex, rightSellDateIndex
	}
}

func shuffleArray2(A []int, left, right int) {
	c := left + (right-left)/2
	q := 1 + left + (c-left)/2
	if left == right {
		return
	}

	k, i := 1, q
	for i <= c {
		A[i], A[c+k] = A[c+k], A[i]
		i++
		k++
	}
	shuffleArray2(A, left, c)
	shuffleArray2(A, c+1, right)
}

func maxSubArray(A []int) int {
	length := len(A)
	low := 0
	high := length - 1
	_, _, sum := findMaxSubArray(A, low, high)
	return sum
}

func findMaxSubArray(A []int, low int, high int) (int, int, int) {
	if high == low {
		return low, high, A[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := findMaxSubArray(A, low, mid)
		rightLow, rightHigh, rightSum := findMaxSubArray(A, mid+1, high)
		crossLow, crossHigh, crossSum := maxCrossingSubArray(A, low, high, mid)

		if (leftSum >= rightSum) && (leftSum >= crossSum) {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

func maxCrossingSubArray(A []int, low, high, mid int) (int, int, int) {
	leftSum := math.MinInt32
	rightSum := math.MinInt32
	maxLeft, maxRight, sum := 0, 0, 0
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, (leftSum + rightSum)
}
