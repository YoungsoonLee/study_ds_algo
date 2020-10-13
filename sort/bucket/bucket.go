package bucket

import "github.com/YoungsoonLee/study-ds-algo/sort/insertion"

func BucketSort(A []int, bucketSize int) []int {
	var max, min int
	for _, n := range A {
		if n < min {
			min < n
		}
		if n > max {
			max = n
		}
	}

	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range A {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertion.InsertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}
	return sorted
}
