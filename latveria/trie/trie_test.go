package trie

import (
	"fmt"
	"github.com/Bishoptylaor/go-toolkit/zrand"
	"sync"
	"testing"
)

func TestSafeTrie(t *testing.T) {
	st := NewSafeTrie[int]()
	wg1, wg2 := sync.WaitGroup{}, sync.WaitGroup{}
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	go func() {
		select {
		case <-ch1:
			for i := 0; i < 100; i++ {
				wg1.Add(1)
				go func() {
					defer wg1.Done()
					word := zrand.RandString(10)
					st.Insert(word, 0)
					fmt.Println(word)
				}()
			}
		}
	}()

	go func() {
		select {
		case <-ch2:
			for i := 0; i < 100; i++ {
				wg2.Add(1)
				go func() {
					defer wg2.Done()
					word := zrand.RandString(8)
					res := st.Search(word)
					fmt.Println(word, res)
				}()
			}
		}
	}()

	ch1 <- struct{}{}
	ch2 <- struct{}{}
	wg1.Wait()
	wg2.Wait()
	fmt.Println("end")
}

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

func TestPath(t *testing.T) {
	path := NewPathTrie()
	path.Insert("/a/b/c")
	pp(path.root, 1)
	fmt.Println("--------------------")
	path.Insert("/a/d/e")
	pp(path.root, 1)
	fmt.Println("--------------------")
	path.Insert("/x/y")
	pp(path.root, 1)
	fmt.Println("--------------------")
	path.Insert("/z")
	pp(path.root, 1)
	fmt.Println("--------------------")
	path.Erase("/z")
	pp(path.root, 1)
	fmt.Println("--------------------")
	path.Erase("/a/b")
	pp(path.root, 1)
	fmt.Println("--------------------")
	path.Erase("/a/b/c")
	pp(path.root, 1)
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

func pp(node *pathTrieNode, level int) {
	if node == nil {
		return
	}
	fmt.Println("current", node.part, "level", level, "passcnt:", node.passCnt, "end:", node.end)
	for c, _ := range node.children {
		fmt.Println("c:", string(c))
	}
	for _, child := range node.children {
		pp(child, level+1)
	}
}
