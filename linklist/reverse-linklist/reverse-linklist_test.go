package reverse_linklist

import (
	"testing"
)

func Test_reverseLinkList(t *testing.T) {
	head := &LinkNode{1, &LinkNode{2, &LinkNode{3, &LinkNode{4, nil}}}}
	result := reverseLinkList(head)
	result.Dump()
}

func Test_reverseLinkList2(t *testing.T) {
	head := &LinkNode{1, &LinkNode{2, &LinkNode{3, &LinkNode{4, nil}}}}
	result := reverseLinkList2(head)
	result.Dump()
}
