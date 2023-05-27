package ds

import "container/heap"

// priorityQueueListは優先度付きキューのリストを表す
// この型はPriorityQueueとheapが要求するメソッド名が重複しないようにするための
type priorityQueueList struct {
	values []interface{}
	prior  func(a, b interface{}) bool
}

// PriorityQueue は優先度付きキューを表す
type PriorityQueue struct {
	list *priorityQueueList
}

// NewPriorityQueueはpriorで優先度が決まる空の優先度付きキューを返します。
func NewPriorityQueue(prior func(a, b interface{}) bool) *PriorityQueue {
	return &PriorityQueue{
		list: &priorityQueueList{
			values: make([]interface{}, 0),
			prior:  prior,
		},
	}
}

// Push は優先度付きキューに要素を一つ追加します。
func (pq PriorityQueue) Push(value interface{}) {
	heap.Push(pq.list, value)
}

// Pop は優先度付きキューから要素を一つ取り出します。
func (pq PriorityQueue) Pop() interface{} {
	return heap.Pop(pq.list)
}

// Top は優先度つきキューの先頭要素を返します。
func (pq PriorityQueue) Top() interface{} {
	v := heap.Pop(pq.list)
	heap.Push(pq.list, v)
	return v
}

// Empty は優先度付きキューが空かどうかを判断します。
func (pq PriorityQueue) Empty() bool {
	return pq.list.Len() == 0
}

// HasElementsはこの優先度付きキューが要素を持つかどうかを判断します。
func (pq PriorityQueue) HasElements() bool {
	return !pq.Empty()
}

// Swap は要素を交換します。
func (list priorityQueueList) Swap(i, j int) {
	list.values[i], list.values[j] = list.values[j], list.values[i]
}

// Less は要素を比較し、優先度が低いかどうかを判断します
func (list priorityQueueList) Less(i, j int) bool {
	return list.prior(list.values[i], list.values[j])
}

// Len は要素の数を返します。
func (list priorityQueueList) Len() int {
	return len(list.values)
}

// Pop は要素を取り出して返します。
func (list *priorityQueueList) Pop() interface{} {
	old := list.values
	n := len(old)
	item := old[n-1]
	values := old[:n-1]
	list.values = values
	return item
}

// Push は要素を追加します。
func (list *priorityQueueList) Push(item interface{}) {
	list.values = append(list.values, item)
}
