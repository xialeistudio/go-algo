package countsort

import "testing"

func Test_countsort(t *testing.T) {
	nums := []int{9, 9, 10, 8, 7, 6, 5, 4, 3, 2, 1}
	countsort(nums)
	t.Log(nums)
}
