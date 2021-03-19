package maxup_subarray

import "testing"

func Test_maxUpSubArray(t *testing.T) {
	nums := []int{10, 9, 2, 10, 9, 5, 7, 101, 18}
	t.Log(maxUpSubArray(nums))
}
