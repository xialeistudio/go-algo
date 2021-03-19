package min_triange

import "testing"

func Test_minTriange(t *testing.T) {
	nums := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	t.Log(minTriange(nums))
}
