package ds

import rbt "github.com/emirpasic/gods/trees/redblacktree"

type Map[K, V any] struct {
	Put          func(key K, value V)
	Size         func() int
	Empty        func() bool
	Remove       func(key K) bool
	Get          func(key K) (V, bool)
	GetOrDefault func(key K, defaultValue V) V
	Left         func() K
	Right        func() K
	Floor        func(key K) (K, V, bool)
	Ceiling      func(key K) (K, V, bool)
}

func NewMap[K, V any](compare func(a, b K) int) *Map[K, V] {
	tree := rbt.NewWith(func(a, b any) int {
		return compare(a.(K), b.(K))
	})

	return &Map[K, V]{
		Put: func(key K, value V) {
			tree.Put(key, value)
		},
		Size: func() int {
			return tree.Size()
		},
		Empty: func() bool {
			return tree.Empty()
		},
		Remove: func(key K) bool {
			if _, ex := tree.Get(key); ex {
				tree.Remove(key)
				return true
			} else {
				return false
			}
		},
		Get: func(key K) (V, bool) {
			if value, ok := tree.Get(key); ok {
				return value.(V), true
			}
			var zero V
			return zero, false
		},
		GetOrDefault: func(key K, defaultValue V) V {
			if value, ok := tree.Get(key); ok {
				return value.(V)
			}
			return defaultValue
		},
		Left: func() K {
			node := tree.Left()
			if node == nil {
				panic("Left() called on empty tree")
			}
			return node.Key.(K)
		},
		Right: func() K {
			node := tree.Right()
			if node == nil {
				panic("Right() called on empty tree")
			}
			return node.Key.(K)
		},
		Floor: func(key K) (K, V, bool) {
			if node, ex := tree.Floor(key); ex {
				return node.Key.(K), node.Value.(V), true
			}
			var zero K
			var zeroV V
			return zero, zeroV, false
		},
		Ceiling: func(key K) (K, V, bool) {
			if node, ex := tree.Ceiling(key); ex {
				return node.Key.(K), node.Value.(V), true
			}
			var zero K
			var zeroV V
			return zero, zeroV, false
		},
	}
}

func NewIntMap[V any]() *Map[int, V] {
	return NewMap[int, V](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
}
