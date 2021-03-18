package quicksort

import "testing"

func Test_quicksort(t *testing.T) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quicksort(nums, 0, len(nums)-1)
	t.Log(nums)
}
