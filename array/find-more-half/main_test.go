package find_more_half

import "testing"

func Test_findMoreHalf(t *testing.T) {
	nums := []int{1, 2, 1, 1, 3, 1, 4, 1, 1, 2, 1, 3, 1, 4, 1, 5, 1}
	t.Log(findMoreHalf(nums))
}
