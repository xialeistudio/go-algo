package multimerge

import (
	"testing"
)

func Test_multiMerge(t *testing.T) {
	nums := [][]int{
		{9, 10, 11, 12},
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	result := multimerge(nums)
	t.Log(result)
}
