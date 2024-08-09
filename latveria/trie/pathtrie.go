package trie

import (
	"strings"
)

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
 @Time    : 2024/8/06 -- 14:07
 @Author  : bishop ❤️ MONEY
 @Description: 路径 trie 树，专为 uri path ，根据 /a /bc /d 分割节点
*/

type PathTrie struct {
	root      *pathTrieNode
	segmenter StringSegmenter
}

type pathTrieNode struct {
	children map[string]*pathTrieNode // children
	part     string                   // current part of fullPath
	passCnt  int                      // counts of how many word has same node; including itself
	end      bool                     // if current node is the end of a word
}

// StringSegmenter 定义分割字符串方法，方便自定义分割方案
type StringSegmenter func(key string, start int) (segment string, nextIndex int)

func PathSegmenter(path string, start int) (segment string, next int) {
	if len(path) == 0 || start < 0 || start > len(path)-1 {
		return "", -1
	}
	end := strings.IndexRune(path[start+1:], '/') // next '/' after 0th rune
	if end == -1 {
		return path[start:], -1
	}
	return path[start : start+end+1], start + end + 1
}

func NewPathTrie() *PathTrie {
	return &PathTrie{
		root: &pathTrieNode{
			children: make(map[string]*pathTrieNode),
		},
		segmenter: PathSegmenter,
	}
}

func (t *PathTrie) Insert(word string) {
	// todo check if word matches path pattern
	if t.Search(word) {
		return
	}

	ptr := t.root
	for part, i := t.segmenter(word, 0); part != ""; part, i = t.segmenter(word, i) {
		if _, ok := ptr.children[part]; !ok {
			// init
			ptr.children[part] = &pathTrieNode{children: make(map[string]*pathTrieNode)}
		}
		// cnt ++
		ptr.children[part].passCnt++
		ptr.children[part].part = part
		// go to next ch
		ptr = ptr.children[part]
	}

	ptr.end = true
}

func (t *PathTrie) Search(word string) bool {
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

func (t *PathTrie) searchWord(word string) *pathTrieNode {
	ptr := t.root
	for part, i := t.segmenter(word, 0); part != ""; part, i = t.segmenter(word, i) {
		if ptr.children[part] == nil {
			return nil
		}
		ptr = ptr.children[part]
	}
	return ptr
}

func (t *PathTrie) StartsWith(prefix string) bool {
	return t.searchWord(prefix) != nil
}

func (t *PathTrie) PassCnt(prefix string) int {
	node := t.searchWord(prefix)
	if node == nil {
		return 0
	}
	return node.passCnt
}

func (t *PathTrie) Erase(word string) bool {
	if !t.Search(word) {
		return false
	}

	ptr := t.root
	for part, i := t.segmenter(word, 0); part != ""; part, i = t.segmenter(word, i) {
		ptr.children[part].passCnt--
		// passCnt == 0 = this is the end of this word.
		if ptr.children[part].passCnt == 0 {
			delete(ptr.children, part)
			return true
		}
		ptr = ptr.children[part]
	}
	// if there is still more ch in this branch. should remove the 'end' tag
	ptr.end = false
	return true
}
