// 三角形最小路径和
package min_triange

func minTriange(nums [][]int) int {
	dp := make([][]int, len(nums))
	for i, row := range nums {
		dp[i] = make([]int, len(row))
	}
	// 初始状态
	dp[0][0] = nums[0][0]
	// 开始处理
	for i := 1; i < len(nums); i++ {
		// 行
		row := nums[i]
		for j := 0; j < len(row); j++ {
			// 直接下来的
			if j == 0 {
				dp[i][j] = dp[i-1][j] + nums[i][j]
				continue
			}
			// 正上方为空的，只能右下角
			if len(dp[i-1]) == j {
				dp[i][j] = dp[i-1][j-1] + nums[i][j]
				continue
			}
			// 比较直接下来和右下角下来哪个小，小的作为当前位置的最小路径值
			dp[i][j] = min(dp[i-1][j]+nums[i][j], dp[i-1][j-1]+nums[i][j])
		}
	}
	// 遍历最后一行结果，找到最小值
	result := dp[len(dp)-1][0]
	for _, value := range dp[len(dp)-1] {
		result = min(value, result)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
