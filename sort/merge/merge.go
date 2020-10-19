package merge

func MergeSort(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	mid := len(A) / 2
	left := MergeSort(A[:mid])
	right := MergeSort(A[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			// 나머지 처리
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			// 나머지 처리
			result[i] = right[0]
			right = right[1:]
		}

	}
	return result
}

func MergeSortTucker(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	left := MergeSortTucker(arr[:len(arr)/2])
	right := MergeSortTucker(arr[len(arr)/2:])

	i, j := 0, 0
	rst := make([]int, 0)
	for i < len(left) || j < len(right) {
		if i >= len(left) {
			rst = append(rst, right[j])
			j++
		} else if j >= len(right) {
			rst = append(rst, left[i])
			i++
		} else {
			if left[i] < right[j] {
				rst = append(rst, left[i])
				i++
			} else {
				rst = append(rst, left[j])
				j++
			}
		}
	}
	return rst
}
