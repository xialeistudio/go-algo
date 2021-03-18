// 计数排序
package countsort

func countsort(nums []int) {
	max := 0
	// 获取数组最大值
	for _, val := range nums {
		if val > max {
			max = val
		}
	}

	// 类似map,key为数组的元素，值为出现次数
	tmp := make([]int, max+1)
	for _, val := range nums {
		tmp[val]++
	}
	j := 0
	for i := 0; i < max+1; i++ {
		// tmp数组不为空的情况下，证明当前位置有值，需要插入最终结果
		for tmp[i] > 0 {
			nums[j] = i
			j++
			tmp[i]--
		}
	}
}
