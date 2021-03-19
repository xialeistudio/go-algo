package reverse_linklist

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

// 迭代法
func reverseLinkList(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		prev    *LinkNode
		next    *LinkNode
		current = head
	)
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

// 递归
func reverseLinkList2(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseLinkList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
