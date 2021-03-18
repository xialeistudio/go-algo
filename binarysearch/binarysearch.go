// 二分查找
package binarysearch

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		value := nums[mid]
		// 中间值小于目标值，则目标值在右边
		if value < target {
			left = mid + 1
			continue
		}
		// 中间值大于目标值，则目标值在左边
		if value > target {
			right = mid - 1
			continue
		}
		// 查找成功
		return mid
	}
	return -1
}
