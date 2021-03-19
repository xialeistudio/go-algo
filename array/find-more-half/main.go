package find_more_half

func findMoreHalf(nums []int) int {
	value := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == value {
			count++
			continue
		}
		count--
		if count < 0 {
			value = nums[i]
			count = 1
		}
	}
	return value
}
