package ds

import (
	"bytes"
	"fmt"
)

type LargeBitSetNode struct {
	// [left, right)を管理します。
	left, right int

	// childrenのうち値をもつ要素の数
	used int

	// このセットの子ノード. 64分木の形で実装する
	children [64]*LargeBitSetNode
}

func (n *LargeBitSetNode) String() string {
	exists := bytes.Repeat([]byte{'.'}, 64)
	for i, v := range n.children {
		if v != nil {
			exists[i] = 'o'
		}
	}

	return fmt.Sprintf("node[%d, %d] %d: %s", n.left, n.right, n.used, exists)
}

// Rangeはこのノードが持つ区間集合のサイズを返します。
func (n *LargeBitSetNode) Range() int {
	return (n.right - n.left) / 64
}

// Indexはこのノードが持つ子要素のうち、vに対応するインデックスを返します。
func (n *LargeBitSetNode) Index(v int) int {
	return (v - n.left) / n.Range()
}

// 巨大サイズの非負整数の集合を管理します。
type LargeBitSet struct {
	// ルートノード
	root *LargeBitSetNode

	// Addはこの集合にvを追加します。
	Add func(v int)

	// Hasはこの集合にvが含まれるかどうかを判断します。
	Has func(v int) bool

	// Removeはこの集合からvを削除します。
	Remove func(v int)

	// Beforeはこの集合に含まれるv未満の要素のうち、最大の要素を返します。
	// 存在しない場合はnilを返します。
	Before func(v int) *int

	// Afterはこの集合に含まれるvを超える要素のうち、最小の要素を返します。
	// 存在しない場合はnilを返します。
	After func(v int) *int
}

func NewLargeBitSet() *LargeBitSet {
	set := LargeBitSet{
		root: &LargeBitSetNode{
			left:     0,
			right:    1 << 60, // 64の累乗の形であらわすと1<<60が限界になる
			children: [64]*LargeBitSetNode{},
		},
	}
	set.Add = func(v int) {
		p := set.root
		for p.Range() > 0 {
			i := p.Index(v)
			if p.children[i] == nil {
				child := LargeBitSetNode{
					left:  p.left + (i+0)*p.Range(),
					right: p.left + (i+1)*p.Range(),
				}
				if p.Range() > 1 {
					child.children = [64]*LargeBitSetNode{}
				}
				p.used++
				p.children[i] = &child
			}
			p = p.children[i]
		}
	}

	set.Has = func(v int) bool {
		p := set.root
		for p != nil {
			i := p.Index(v)
			if p.Range() == 1 {
				return p.children[i] != nil
			}
			p = p.children[i]
		}
		return false
	}

	var remove func(n *LargeBitSetNode, v int)
	remove = func(n *LargeBitSetNode, v int) {
		i := n.Index(v)
		if n.children[i] == nil {
			return
		}

		if n.Range() > 1 {
			remove(n.children[i], v)
		}
		if n.children[i].used == 0 {
			n.children[i] = nil
			n.used--
		}
	}

	set.Remove = func(v int) {
		remove(set.root, v)
	}

	var before func(n *LargeBitSetNode, v int) *int
	before = func(n *LargeBitSetNode, v int) *int {
		for i := n.Index(v); i >= 0; i-- {
			if n.children[i] == nil {
				continue
			}
			if n.Range() == 1 {
				ret := n.left + i
				return &ret
			}
			ret := before(n.children[i], v)
			if ret != nil {
				return ret
			}
		}
		return nil
	}
	set.Before = func(v int) *int {
		return before(set.root, v-1)
	}

	var after func(n *LargeBitSetNode, v int) *int
	after = func(n *LargeBitSetNode, v int) *int {
		for i := n.Index(v); i < 64; i++ {
			if n.children[i] == nil {
				continue
			}
			if n.Range() == 1 {
				ret := n.left + i
				return &ret
			}

			ret := after(n.children[i], v)
			if ret != nil {
				return ret
			}
		}
		return nil
	}
	set.After = func(v int) *int {
		return after(set.root, v+1)
	}

	return &set
}
