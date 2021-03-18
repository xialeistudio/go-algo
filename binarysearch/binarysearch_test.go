package binarysearch

import "testing"

func Test_binarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := binarySearch(nums, 10)
	t.Log(index)
}
