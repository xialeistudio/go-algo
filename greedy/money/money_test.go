package money

import (
	"testing"
)

func Test_money(t *testing.T) {
	money := 77
	values := []int{25, 10, 5, 1}
	result := takeMoney(money, values)
	t.Log(result)
}
