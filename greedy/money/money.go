// 贪心算法：找零钱
package money

// takeMoney 总钱数
// values 排好序的钱币，由大到小
func takeMoney(money int, values []int) map[int]int {
	// 有25、10、5、1可以使用
	result := make(map[int]int)
	for _, value := range values {
		for money >= value {
			result[value]++
			money -= value
		}
	}
	return result
}
