// 背包问题
package pack

import "fmt"

type status int

const (
	statusInit    status = iota // 初始状态
	statusPicked                // 已选择
	statusInvalid               // 不可以选择
)

type item struct {
	weight int
	price  int
	status status
}

// 策略
type strategy interface {
	pick(items []item) int
}

// 重量优先选择算法
type weightPriorityStrategy struct{}

func (w weightPriorityStrategy) pick(items []item) int {
	index := -1
	minWeight := 100
	for i, item := range items {
		if item.status == statusInit && item.weight < minWeight {
			index = i
			minWeight = item.weight
		}
	}
	return index
}

// 价格优先
type pricePriorityStrategy struct{}

func (p pricePriorityStrategy) pick(items []item) int {
	index := -1
	price := 0
	for i, item := range items {
		if item.status == statusInit && item.price > price {
			index = i
			price = item.price
		}
	}
	return index
}

type densityPriorityStrategy struct{}

func (d densityPriorityStrategy) pick(items []item) int {
	var (
		density float32
		index   = -1
	)
	for i, item := range items {
		itemDensity := float32(item.price) / float32(item.weight)
		if item.status == statusInit && itemDensity > density {
			index = i
			density = itemDensity
		}
	}
	return index
}

// 选择入口
// strategy 选择策略
func pick(strategy strategy) {
	items := []item{
		{35, 10, statusInit},
		{30, 40, statusInit},
		{60, 30, statusInit},
		{50, 50, statusInit},
		{40, 35, statusInit},
		{10, 40, statusInit},
		{25, 30, statusInit},
	}
	maxWeight := 150
	totalWeight := 0
	for {
		index := strategy.pick(items)
		// 无可用item
		if index == -1 {
			break
		}
		item := items[index]
		// 装不下了
		if (totalWeight + item.weight) > maxWeight {
			// 标记物品为不符合条件
			items[index].status = statusInvalid
			continue
		}
		items[index].status = statusPicked
		totalWeight += item.weight
	}
	dump(items)
}

// 打印结果
func dump(items []item) {
	totalWeight := 0
	totalPrice := 0
	for index, item := range items {
		if item.status != statusPicked {
			continue
		}
		totalWeight += item.weight
		totalPrice += item.price
		fmt.Printf("index=%v, weight=%v, price=%v density=%v\n", index, item.weight, item.price, float32(item.price)/float32(item.weight))
	}
	fmt.Printf("totalWeight=%v, totalPrice=%v\n", totalWeight, totalPrice)
}
