package pack

import "testing"

// 重量优先选择
func Test_weightPriority(t *testing.T) {
	strategy := &weightPriorityStrategy{}
	pick(strategy)
}

// 价值优先选择
func Test_pricePriority(t *testing.T) {
	strategy := &pricePriorityStrategy{}
	pick(strategy)
}

// 价值密度优先选择
func Test_densityPriority(t *testing.T) {
	strategy := &densityPriorityStrategy{}
	pick(strategy)
}
