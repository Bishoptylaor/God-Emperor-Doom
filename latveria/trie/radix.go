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
 @Description: 基数树（Radix Tree）也称为压缩前缀树（compact prefix tree）或 compressed trie 基础实现
	imitation of https://github.com/gin-gonic/gin/blob/master/tree.go
*/

type Radix struct {
	root *radixNode
}

type radixNode struct {
	path     string       // relative path of this node
	fullPath string       // full path of this node
	indices  string       // aka []byte{children[0].path[0], children[1].path[0], children[2].path[0]}
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

	// erase root
	if r.root.fullPath == word {
		if len(r.root.indices) == 0 {
			r.root.path = ""
			r.root.end = false
			r.root.fullPath = ""
			r.root.passCnt = 0
			return true
		}

		if len(r.root.indices) == 1 {
			r.root.children[0].path = r.root.path + r.root.children[0].path
			r.root = r.root.children[0]
			return true
		}

		for i := 0; i < len(r.root.indices); i++ {
			r.root.children[i].path = r.root.path + r.root.children[i].path
		}

		newRoot := &radixNode{
			indices:  r.root.indices,
			children: r.root.children,
			passCnt:  r.root.passCnt - 1,
		}
		r.root = newRoot
		return true
	}

	ptr := r.root
walk:
	for {
		ptr.passCnt -= 1
		prefix := ptr.path
		word := word[len(prefix):]
		c := word[0]
		for i := 0; i < len(ptr.indices); i++ {
			if ptr.indices[i] != c {
				continue
			}

			if ptr.children[i].path == word && ptr.children[i].passCnt > 1 {
				ptr.children[i].end = false
				ptr.children[i].passCnt -= 1
				return true
			}

			if ptr.children[i].passCnt > 1 {
				ptr = ptr.children[i]
				continue walk
			}

			ptr.children = append(ptr.children[:i], ptr.children[i+1:]...)
			ptr.indices = ptr.indices[i:] + ptr.indices[i+1:]
			if !ptr.end && len(ptr.indices) == 1 {
				ptr.path += ptr.children[0].path
				ptr.fullPath = ptr.children[0].fullPath
				ptr.end = ptr.children[0].end
				ptr.indices = ptr.children[0].indices
				ptr.children = ptr.children[0].children
			}

			return true
		}
	}
}

func (rn *radixNode) insert(word string) {
	fullWord := word

	// first word
	if rn.path == "" && len(rn.children) == 0 {
		rn.insertWord(word, word)
		return
	}

walk:
	for {
		i := longestCommonPrefix(word, rn.path)
		if i > 0 {
			rn.passCnt += 1
		}

		if i < len(rn.path) {
			child := radixNode{
				path:     rn.path[:],
				fullPath: rn.fullPath,
				indices:  rn.indices,
				children: rn.children,
				end:      rn.end,
				passCnt:  rn.passCnt - 1,
			}

			rn.indices = string(rn.path[i])
			rn.children = append(rn.children, &child)
			rn.fullPath = rn.fullPath[:len(rn.fullPath)-len(rn.path)+i]
			rn.path = rn.path[:i]
			rn.end = false
		}

		if i < len(word) {
			word = word[i:]
			w := word[0]
			for i := 0; i < len(rn.indices); i++ {
				if rn.indices[i] == w {
					rn = rn.children[i]
					continue walk
				}
			}

			rn.indices += string(w)
			child := radixNode{}
			child.insertWord(word, fullWord)
			rn.children = append(rn.children, &child)
			return
		}

		rn.end = true
		return
	}
}

func (rn *radixNode) insertWord(path, fullPath string) {
	rn.fullPath = fullPath
	rn.path = path
	rn.passCnt = 1
	rn.end = true
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max_ := min(len(a), len(b))
	for i < max_ && a[i] == b[i] {
		i++
	}
	return i
}

func (rn *radixNode) search(word string) *radixNode {
walk:
	for {
		prefix := rn.path
		// word longer then current node prefix
		if len(word) > len(prefix) {
			if word[:len(prefix)] != prefix {
				return nil
			}
			word = word[len(prefix):]
			c := word[0]
			for i := 0; i < len(rn.indices); i++ {
				if c == rn.indices[i] {
					rn = rn.children[i]
					continue walk
				}
			}

			return nil
		}

		if word == prefix {
			return rn
		}

		return rn
	}
}
