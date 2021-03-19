// 爬楼梯
package floor

// 爬楼梯
// n楼梯阶数
// 返回爬楼梯方法
func floor(n int) int {
	if n <= 2 {
		return n
	}
	result := 0
	one := 1
	two := 2
	for i := 3; i <= n; i++ {
		// 状态转移方程
		result = one + two
		one = two
		two = result
	}
	return result
}
