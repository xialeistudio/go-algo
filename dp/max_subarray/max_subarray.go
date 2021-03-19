// 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
package max_subarray

// 最大子序和
// nums 数字序列
// @return 最大子序列和
// 分析
// 记位置i的最大子序和为dp[i]
// 那么有 dp[i] = dp[i-1] + nums[i]
// 如果dp[i-1] 小于0，则dp[i] = nums[i]
// 已知dp[0] = nums[0]
func maxSubArray(nums []int) int {
	pre := 0
	result := nums[0]
	for i, num := range nums {
		if i == 0 {
			pre = num
			continue
		}
		pre = max(pre+num, num)
		result = max(pre, result)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
