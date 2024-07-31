package trie

import (
	"fmt"
	"testing"
)

func TestRadix(t *testing.T) {
	radix := NewRadix()
	radix.Insert("hello world")
	pr(radix.root, 1)
	fmt.Println("--------------------")
	radix.Insert("hello leetcode")
	pr(radix.root, 1)
	fmt.Println("--------------------")
	radix.Insert("hello leet678")
	pr(radix.root, 1)
	fmt.Println("--------------------")
	radix.Insert("hello foo")
	pr(radix.root, 1)
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	//trie.Insert("hello world")
	//pt(trie.root, 1)
	//fmt.Println("--------------------")
	trie.Insert("red")
	pt(trie.root, 1)
	/*
		level 1 passcnt: 0 end: false
		c: r
		level 2 passcnt: 1 end: false
		c: e
		level 3 passcnt: 1 end: false
		c: d
		level 4 passcnt: 1 end: true
	*/
	fmt.Println("--------------------")

	trie.Insert("redlock")
	pt(trie.root, 1)
	/*
		level 1 passcnt: 0 end: false
		c: r
		level 2 passcnt: 2 end: false
		c: e
		level 3 passcnt: 2 end: false
		c: d
		level 4 passcnt: 2 end: true
		c: l
		level 5 passcnt: 1 end: false
		c: o
		level 6 passcnt: 1 end: false
		c: c
		level 7 passcnt: 1 end: false
		c: k
		level 8 passcnt: 1 end: true
	*/
	fmt.Println("--------------------")

	trie.Insert("redlocker")
	pt(trie.root, 1)
	/*
		level 1 passcnt: 0 end: false
		c: r
		level 2 passcnt: 3 end: false
		c: e
		level 3 passcnt: 3 end: false
		c: d
		level 4 passcnt: 3 end: true
		c: l
		level 5 passcnt: 2 end: false
		c: o
		level 6 passcnt: 2 end: false
		c: c
		level 7 passcnt: 2 end: false
		c: k
		level 8 passcnt: 2 end: true
		c: e
		level 9 passcnt: 1 end: false
		c: r
		level 10 passcnt: 1 end: true
	*/
	fmt.Println("--------------------")

	trie.Erase("redlocker")
	pt(trie.root, 1)
	/*
		level 1 passcnt: 0 end: false
		c: r
		level 2 passcnt: 2 end: false
		c: e
		level 3 passcnt: 2 end: false
		c: d
		level 4 passcnt: 2 end: true
		c: l
		level 5 passcnt: 1 end: false
		c: o
		level 6 passcnt: 1 end: false
		c: c
		level 7 passcnt: 1 end: false
		c: k
		level 8 passcnt: 1 end: true
	*/
	fmt.Println("--------------------")

}

func pr(node *radixNode, level int) {
	if node == nil {
		return
	}
	fmt.Println("level", level, "fullpath:", node.fullPath, "path:", node.path, "passcnt:", node.passCnt, "end:", node.end, "indices:", node.indices)
	for _, child := range node.children {
		pr(child, level+1)
	}
}

func pt(node *trieNode, level int) {
	if node == nil {
		return
	}
	fmt.Println("level", level, "passcnt:", node.passCnt, "end:", node.end)
	for c, _ := range node.children {
		fmt.Println("c:", string(c))
	}
	for _, child := range node.children {
		pt(child, level+1)
	}
}
