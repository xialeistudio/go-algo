package mergesort

import (
	"testing"
)

func Test_merge(t *testing.T) {
	left := []int{1, 2, 3, 4}
	right := []int{5, 6, 7, 8}
	result := merge(left, right)
	t.Log(result)
}

func Test_sort(t *testing.T) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := sort(nums)
	t.Log(result)
}
