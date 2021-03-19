// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
package maxup_subarray

func maxUpSubArray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // 当前初始化长度1
		// 往前找，有符合条件的数字时则dp[i]继续加
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // 前面的数比当前数字小
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(dp[i], result)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
