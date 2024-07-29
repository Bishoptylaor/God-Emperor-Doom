package triedoll

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
 @Description: trie 树基础实现
*/

type Trie struct {
	root *trieNode
}

type trieNode struct {
	children map[rune]*trieNode
	passCnt  int
	end      bool
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{},
	}
}

func (t *Trie) Insert(word string) {
	if t.Search(word) {
		return
	}

	ptr := t.root
	for _, ch := range word {
		if _, ok := ptr.children[ch]; !ok {
			// init
			ptr.children[ch] = &trieNode{}
		} else {
			// cnt ++
			ptr.children[ch].passCnt++
			// go to next ch
			ptr = ptr.children[ch]
		}
	}

	ptr.end = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchWord(word)
	// 没找到
	if node == nil {
		return false
	}
	// 确实存在单词以该节点作为结尾
	if node.end {
		return true
	}
	return false
}

func (t *Trie) searchWord(word string) *trieNode {
	ptr := t.root
	for _, ch := range word {
		if ptr.children[ch] == nil {
			return nil
		}
		ptr = ptr.children[ch]
	}
	return ptr
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchWord(prefix) != nil
}

func (t *Trie) PassCnt(prefix string) int {
	node := t.searchWord(prefix)
	if node == nil {
		return 0
	}
	return node.passCnt
}

func (t *Trie) Erase(word string) bool {
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
		ptr = ptr.children[ch]
	}
	// if there is still more ch in this branch. should remove the 'end' tag
	ptr.end = false
	return true
}
