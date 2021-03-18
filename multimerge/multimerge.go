// 多路归并
package multimerge

func multimerge(nums [][]int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	if length == 1 {
		return nums[0]
	}
	// 超过两路
	mid := length / 2
	left := multimerge(nums[:mid])
	right := multimerge(nums[mid:])
	return merge(left, right)
}

// 二路有序数组归并
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
