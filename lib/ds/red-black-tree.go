package ds

import "github.com/emirpasic/gods/trees/redblacktree"

type RedBlackTree[K any, V any] struct {
	delegate *redblacktree.Tree
}

func NewRedBlackTree[K any, V any](comparator func(a, b K) int) *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{
		delegate: redblacktree.NewWith(func(x, y interface{}) int {
			return comparator(x.(K), y.(K))
		}),
	}
}

func (tree *RedBlackTree[K, V]) Ceiling(key K) (*K, *V) {
	node, found := tree.delegate.Ceiling(key)
	if !found {
		return nil, nil
	} else {
		key, value := node.Key.(K), node.Value.(V)
		return &key, &value
	}
}

func (tree *RedBlackTree[K, V]) Floor(key K) (*K, *V) {
	node, found := tree.delegate.Floor(key)
	if !found {
		return nil, nil
	} else {
		key, value := node.Key.(K), node.Value.(V)
		return &key, &value
	}
}

func (tree *RedBlackTree[K, V]) Empty() bool {
	return tree.delegate.Empty()
}

func (tree *RedBlackTree[K, V]) HasEntry() bool {
	return !tree.delegate.Empty()
}

func (tree *RedBlackTree[K, V]) Put(key K, value V) {
	tree.delegate.Put(key, value)
}

func (tree *RedBlackTree[K, V]) Get(key K) *V {
	node := tree.delegate.GetNode(key)
	if node == nil {
		return nil
	}
	value := node.Value.(V)
	return &value
}

func (tree *RedBlackTree[K, V]) Remove(key K) {
	tree.delegate.Remove(key)
}

func (tree *RedBlackTree[K, V]) Left() (*K, *V) {
	node := tree.delegate.Left()
	if node == nil {
		return nil, nil
	}
	key, value := node.Key.(K), node.Value.(V)
	return &key, &value
}

func (tree *RedBlackTree[K, V]) Size() int {
	return tree.delegate.Size()
}

func (tree *RedBlackTree[K, V]) String() string {
	return tree.delegate.String()
}
