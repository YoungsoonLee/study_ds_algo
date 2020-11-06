package main

import (
	"fmt"
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

/*
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
*/

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

func recursiveFibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}

func fibonacci(n int) int {
	var fib []int

	if len(fib) == 0 {
		fib = make([]int, n+1)
	}

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if fib[n] != 0 {
		return fib[n]
	}
	fib[n] = fibonacci(n-1) + fibonacci(n-2)
	return fib[n]
}

func fibonacciDP(n int) int {
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func fibonacciFinal(n int) int {
	a, b, sum := 0, 1, 0
	for i := 1; i < n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

var factorials []int

func factorialDP(n int) int {
	if len(factorials) == 0 {
		factorials = make([]int, n+1)
	}
	if n == 0 || n == 1 {
		return 1
	} else if factorials[n] != 0 {
		return factorials[n]
	} else {
		factorials[n] = n * factorialDP(n-1)
		return factorials[n]
	}
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	cache := make([]int, n)
	cache[0], cache[1] = 1, 2
	for i := 2; i < n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n-1]
}

func maxContiguousSum(A []int) int {
	maxSum, n := 0, len(A)
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			currentSum := 0
			for k := i; k <= j; k++ {
				currentSum += A[k]
			}
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}

func maxContiguousSum2(A []int) int {
	maxSum, n := 0, len(A)
	for i := 1; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += A[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}

func maxContiguousSum3(A []int) int {
	n := len(A)
	M := make([]int, n+1)
	maxSum := 0
	if A[0] > 0 {
		M[0] = A[0]
	} else {
		M[0] = 0
	}
	for i := 1; i < n; i++ {
		if M[i-1]+A[i] > 0 {
			M[i] = M[i-1] + A[i]
		} else {
			M[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		if M[i] > maxSum {
			maxSum = M[i]
		}
	}
	return maxSum
}

func maxContiguousSum4(A []int) int {
	sumSofar, sumEndingHere, n := 0, 0, len(A)
	for i := 1; i < n; i++ {
		sumEndingHere = sumEndingHere + A[i]
		if sumEndingHere < 0 {
			sumEndingHere = 0
			continue
		}
		if sumSofar < sumEndingHere {
			sumSofar = sumEndingHere
		}
	}
	return sumSofar
}

func subsetSum(A []int, n, sum int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 && sum != 0 {
		return false
	}
	if A[n-1] > sum {
		return subsetSum(A, n-1, sum)
	}
	return subsetSum(A, n-1, sum) || subsetSum(A, n-1, sum-A[n-1])
}

func main() {
	fmt.Println(fibonacciFinal(10))
}
