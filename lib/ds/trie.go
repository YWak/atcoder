package ds

import (
	"fmt"
	"strings"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	end  int
	prev *TrieNode
	next map[byte]*TrieNode
}

func NewTrie() *Trie {
	trie := &Trie{}
	trie.root = trie.newNode(nil)

	return trie
}

func (t *Trie) newNode(parent *TrieNode) *TrieNode {
	return &TrieNode{
		end:  0,
		prev: parent,
		next: map[byte]*TrieNode{},
	}
}

// Insertは、sをTrieに追加します。
func (t *Trie) Insert(s []byte) {
	node := t.root
	for _, c := range s {
		// idx := t.indexOf(c)
		next := node.next[c]
		if next == nil {
			next = t.newNode(node)
			node.next[c] = next
		}

		node = next
	}
	node.end++
}

func (t *Trie) findPrefixNode(s []byte) *TrieNode {
	node := t.root
	for _, c := range s {
		if node == nil {
			return nil
		}
		node = node.next[c]
	}
	return node
}

// RemovePrefixは、sをprefixとしてもつ文字列を全て削除し、その数を返す
func (t *Trie) RemovePrefix(s []byte) int {
	node := t.findPrefixNode(s)
	if node == nil {
		return 0
	}
	delete(node.prev.next, s[len(s)-1])

	// このノードから、各ノードを削除していく。
	count := 0
	queue := []*TrieNode{node}
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		count += node.end
		for _, next := range node.next {
			queue = append(queue, next)
		}
	}

	return count
}

// IsPrefixOfは、このTrieがもつ文字列のいずれかがsのprefixであるかどうかを返します。
func (t *Trie) IsPrefixOf(s []byte) bool {
	node := t.root
	for _, c := range s {
		// prefixとしてたどり終わったので、解決とする
		if node.end > 0 {
			break
		}

		// 次のノードがない場合、絶対にprefixではない
		next := node.next[c]
		if next == nil {
			return false
		}
		node = next
	}
	return node.end > 0
}

func (t *Trie) String() string {
	strs := []string{}
	var dfs func(node *TrieNode, prefix string)
	dfs = func(node *TrieNode, prefix string) {
		if node == nil {
			return
		}
		strs = append(strs, fmt.Sprintf("\"%s\": %d", prefix, node.end))
		for c, next := range node.next {
			dfs(next, fmt.Sprintf("%s%c", prefix, c))
		}
	}
	dfs(t.root, "")

	return "{" + strings.Join(strs, ", ") + "}"
}
