package ds

import (
	"container/heap"
)

// PriorityQueueListは優先度付きキューのリストを表す
// この型はPriorityQueueとheapが要求するメソッド名が重複しないようにするための
type PriorityQueueList struct {
	values []interface{}
	prior  func(a, b interface{}) bool
}

// PriorityQueueListは heap.Interfaceを満たす.
var _ heap.Interface = &PriorityQueueList{}

// Less は要素を比較し、優先度が低いかどうかを判断します
func (list PriorityQueueList) Less(i, j int) bool {
	return list.prior(list.values[i], list.values[j])
}

// Len は要素の数を返します。
func (list PriorityQueueList) Len() int {
	return len(list.values)
}

// Swap は要素を交換します。
func (list PriorityQueueList) Swap(i, j int) {
	list.values[i], list.values[j] = list.values[j], list.values[i]
}

// Pop は要素を取り出して返します。
func (list *PriorityQueueList) Pop() interface{} {
	old := list.values
	n := len(old)
	item := old[n-1]
	values := old[:n-1]
	list.values = values
	return item
}

// Push は要素を追加します。
func (list *PriorityQueueList) Push(item interface{}) {
	list.values = append(list.values, item)
}

// PriorityQueue は優先度付きキューを表す
type PriorityQueue[T any] struct {
	list PriorityQueueList

	// Pushは優先度付きキューに要素を一つ追加します。
	Push func(value T)

	// Popは優先度付きキューから要素を一つ取り出します。
	Pop func() T

	// Peekは優先度つきキューの先頭要素を返します。
	Peek func() T

	// IsEmptyは優先度付きキューが空かどうかを判断します。
	IsEmpty func() bool

	// HasElementsはこの優先度付きキューが要素を持つかどうかを判断します。
	HasElements func() bool

	// Countはこの優先度付きキューがもつ要素の数を返します。
	Count func() int
}

// NewPriorityQueueはpriorで優先度が決まる空の優先度付きキューを返します。
func NewPriorityQueue[T any](prior func(a, b T) bool) *PriorityQueue[T] {
	list := PriorityQueueList{
		prior: func(a, b interface{}) bool { return prior(a.(T), b.(T)) },
	}
	return &PriorityQueue[T]{
		list: list,
		Push: func(value T) {
			heap.Push(&list, value)
		},
		Pop: func() T {
			return heap.Pop(&list).(T)
		},
		Peek: func() T {
			return list.values[0].(T)
		},
		IsEmpty: func() bool {
			return list.Len() == 0
		},
		HasElements: func() bool {
			return list.Len() > 0
		},
		Count: func() int {
			return list.Len()
		},
	}
}

// // NewAscendingPriorityQueueは順序が決まっている要素について、昇順に扱う優先度付きキューを作成して返します。
// func NewAscendingPriorityQueue[T constraints.Ordered]() *PriorityQueue[T] {
// 	return NewPriorityQueue(func(a, b T) bool { return a < b })
// }

// // NewDescendingPriorityQueueは順序が決まっている要素について、降順に扱う優先度付きキューを作成して返します。
// func NewDescendingPriorityQueue[T constraints.Ordered]() *PriorityQueue[T] {
// 	return NewPriorityQueue(func(a, b T) bool { return a > b })
// }
