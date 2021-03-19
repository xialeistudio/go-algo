package min_matrix

import "testing"

func Test_minMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	t.Log(minMatrix(matrix))
}
