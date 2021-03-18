// 归并排序
package mergesort

// 递归版本
func sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	// 分割为左右两部分
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	return merge(sort(left), sort(right))
}

// 两路合并
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	resultIndex := 0
	leftIndex := 0
	rightIndex := 0
	// 两个数组公共部分
	for leftIndex < len(left) && rightIndex < len(right) {
		// left小于或等于右边，则left的数字加入最终结果，同时移动指针
		if left[leftIndex] <= right[rightIndex] {
			result[resultIndex] = left[leftIndex]
			resultIndex++
			leftIndex++
			continue
		}
		result[resultIndex] = right[rightIndex]
		resultIndex++
		rightIndex++
	}
	// 合并左边剩余元素
	for leftIndex < len(left) {
		result[resultIndex] = left[leftIndex]
		resultIndex++
		leftIndex++
	}
	// 合并右边剩余元素
	for rightIndex < len(right) {
		result[resultIndex] = right[rightIndex]
		resultIndex++
		rightIndex++
	}
	return result
}
