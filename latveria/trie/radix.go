package trie

import "strings"

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
 @Description: 压缩前缀树||基数树基础实现
*/

type Radix struct {
	root *radixNode
}

type radixNode struct {
	path     string // relative path of this node
	fullPath string // full path of this node
	// 每个 indice 字符对应一个孩子节点的 path 首字母
	indices  string
	children []*radixNode // children
	end      bool         // if current node is the end of a word
	passCnt  int          // counts of how many word has passed this node
}

func NewRadix() *Radix {
	return &Radix{
		root: &radixNode{},
	}
}

func (r *Radix) Insert(word string) {
	if r.Search(word) {
		return
	}
	r.root.insert(word)
}

func (r *Radix) Search(word string) bool {
	node := r.root.search(word)
	return node != nil && node.fullPath == word && node.end
}

func (r *Radix) StartWith(prefix string) bool {
	node := r.root.search(prefix)
	return node != nil && strings.HasPrefix(node.fullPath, prefix)
}

func (r *Radix) PassCnt(prefix string) int {
	node := r.root.search(prefix)
	if node == nil || !strings.HasPrefix(node.fullPath, prefix) {
		return 0
	}
	return node.passCnt
}

func (r *Radix) Erase(word string) bool {
	if !r.Search(word) {
		return false
	}

}

func (r *radixNode) insert(word string) {

}

func (r *radixNode) search(word string) *radixNode {
	return nil
}
