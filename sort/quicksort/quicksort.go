// 快速排序
package quicksort

func quicksort(nums []int, begin, end int) {
	if begin >= end {
		return
	}
	index := partition(nums, begin, end)
	// 左边
	quicksort(nums, begin, index-1)
	// 右边
	quicksort(nums, index+1, end)
}

func partition(nums []int, begin, end int) int {
	i := begin
	j := end
	// 标志元素
	base := nums[begin]

	// 两指针相遇为结束条件
	for i != j {
		// 右指针往左找一个比base小的数
		for nums[j] >= base && i < j {
			j--
		}
		// 左指针往右找一个比base小的数
		for nums[i] <= base && i < j {
			i++
		}
		// 大小交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将标志元素归位
	nums[begin], nums[i] = nums[i], nums[begin]
	return i
}
