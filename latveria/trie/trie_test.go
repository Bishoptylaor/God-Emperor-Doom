package trie

import (
	"fmt"
	"testing"
)

func TestRadix(t *testing.T) {
	radix := NewRadix()
	radix.Insert("hello world")
	p(radix.root, 1)
	fmt.Println("--------------------")
	radix.Insert("hello leetcode")
	p(radix.root, 1)
	fmt.Println("--------------------")
	radix.Insert("hello leet678")
	p(radix.root, 1)
	fmt.Println("--------------------")
	radix.Insert("hello foo")
	p(radix.root, 1)
}

func p(node *radixNode, level int) {
	if node == nil {
		return
	}
	fmt.Println("level", level, "fullpath:", node.fullPath, "path:", node.path, "passcnt:", node.passCnt, "end:", node.end, "indices:", node.indices)
	for _, child := range node.children {
		p(child, level+1)
	}
}
