// O(1)复杂度删除单链表节点
package delete_node

import "fmt"

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func (p *LinkNode) Dump() {
	tmp := p
	for tmp != nil {
		fmt.Println(tmp.Value)
		tmp = tmp.Next
	}
}

func deleteNode(node *LinkNode) {
	if node == nil || node.Next == nil { // 空节点和尾结点不处理
		return
	}
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}
