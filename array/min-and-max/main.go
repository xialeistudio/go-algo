package min_and_max

func minAndMax(nums []int, begin, end int) (int, int) {
	// 只有一个元素
	if begin == end {
		return nums[begin], nums[begin]
	}
	// 两个元素
	if begin+1 == end {
		if nums[begin] < nums[end] {
			return nums[begin], nums[end]
		}
		return nums[end], nums[begin]
	}
	// 二分
	mid := (begin + end) / 2
	// 递归左边
	lmin, lmax := minAndMax(nums, begin, mid)
	// 递归右边
	rmin, rmax := minAndMax(nums, mid, end)

	maxValue := max(lmax, rmax)
	minValue := min(lmin, rmin)
	return minValue, maxValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
