package trie

import "sync"

/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/7/29 -- 14:34
 @Author  : bishop ❤️ MONEY
 @Description: trie 树||前缀树基础实现 + 并发安全
*/

type SafeTrie[T any] struct {
	root *safeTrieNode[T]
	sync.RWMutex
}

type safeTrieNode[T any] struct {
	children map[rune]*safeTrieNode[T] // children
	passCnt  int                       // counts of how many word has same node; including itself
	end      bool                      // if current node is the end of a word
	meta     T
}

func NewSafeTrie[T any]() *SafeTrie[T] {
	return &SafeTrie[T]{
		root: &safeTrieNode[T]{
			children: make(map[rune]*safeTrieNode[T]),
		},
	}
}

func (t *SafeTrie[T]) Insert(word string, meta T) {
	t.Lock()
	defer t.Unlock()
	if t.Search(word) {
		return
	}

	ptr := t.root
	for _, ch := range word {
		if _, ok := ptr.children[ch]; !ok {
			// init
			ptr.children[ch] = &safeTrieNode[T]{children: make(map[rune]*safeTrieNode[T])}
		}
		// cnt ++
		ptr.children[ch].passCnt++
		ptr.children[ch].meta = meta
		// go to next ch
		nextPtr := ptr.children[ch]
		ptr = nextPtr
	}

	ptr.end = true
}

func (t *SafeTrie[T]) Search(word string) bool {
	t.RLock()
	defer t.RUnlock()
	node := t.searchWord(word)
	// no found
	if node == nil {
		return false
	}
	// there is a _word end at this node while word = _word
	if node.end {
		return true
	}
	return false
}

func (t *SafeTrie[T]) searchWord(word string) *safeTrieNode[T] {
	ptr := t.root
	for _, ch := range word {
		if ptr.children[ch] == nil {
			return nil
		}
		nextPtr := ptr.children[ch]
		ptr = nextPtr
	}
	return ptr
}

func (t *SafeTrie[T]) StartsWith(prefix string) bool {
	t.RLock()
	defer t.RUnlock()
	return t.searchWord(prefix) != nil
}

func (t *SafeTrie[T]) PassCnt(prefix string) int {
	t.RLock()
	defer t.RUnlock()
	node := t.searchWord(prefix)
	if node == nil {
		return 0
	}
	return node.passCnt
}

func (t *SafeTrie[T]) Erase(word string) bool {
	t.RLock()
	defer t.RUnlock()
	if !t.Search(word) {
		return false
	}

	ptr := t.root
	for _, ch := range word {
		ptr.children[ch].passCnt--
		// passCnt == 0 = this is the end of this word.
		if ptr.children[ch].passCnt == 0 {
			delete(ptr.children, ch)
			return true
		}
		nextPtr := ptr.children[ch]
		ptr = nextPtr
	}
	// if there is still more ch in this branch. should remove the 'end' tag
	ptr.end = false
	return true
}
