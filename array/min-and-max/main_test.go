package min_and_max

import "testing"

func Test_minAndMax(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(minAndMax(nums, 0, len(nums)-1))
}
