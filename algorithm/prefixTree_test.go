package algorithm

import (
	"fmt"
	"testing"
)

// 实现一棵前缀树
// 前缀树应用的场景是什么呢？
// 类比与链表，节点里面包含数据域与指针域, 那么前缀树的节点是由什么构成的呢？
// 前缀树存的是哪个字符，由这个位置指针是否为空决定，如果为空，表示不存在对应字符，如果不为空，表示存在该字符
type TrieNode struct {
	isWord   bool
	children [26]*TrieNode // 这里使用数组，避免初始化的问题
}

type Trie struct {
	root *TrieNode // 根节点
}

func (t *Trie) insert(word string) {
	cur := t.root

	// golang中字符串注意事项，字符串与字符？
	for _, s := range word {
		diff := s - 'a'
		if cur.children[diff] == nil {
			cur.children[diff] = &TrieNode{}
		}

		cur = cur.children[diff]
	}

	cur.isWord = true
	return
}

// 搜索是否存在前缀树呢
func (t *Trie) search(pre string) bool {
	cur := t.root
	for _, s := range pre {
		diff := s - 'a'
		if cur.children[diff] == nil {
			return false
		}

		cur = cur.children[diff]
	}

	return true
}

func TestTrie(t *testing.T) {
	trie := &Trie{
		root: &TrieNode{},
	}

	trie.insert("hello")
	trie.insert("he")

	fmt.Println(trie.search("he"))
	fmt.Println(trie.search("heo"))
}
