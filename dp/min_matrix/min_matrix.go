// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。说明：每次只能向下或者向右移动一步。
package min_matrix

func minMatrix(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, len(grid))
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 { // 第一行，只能从左边下来
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 { // 第一列，只能从正上方下来
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			// 比较正上方和正左方的路径值，取小的
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
