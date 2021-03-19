package delete_node

import "testing"

func Test_deleteNode(t *testing.T) {
	head := &LinkNode{1, &LinkNode{2, &LinkNode{3, &LinkNode{4, nil}}}}
	deleteNode(head.Next.Next)
	head.Dump()
}
