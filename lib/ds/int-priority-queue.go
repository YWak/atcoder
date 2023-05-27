package ds

type IntPriorityQueue struct {
	queue *PriorityQueue
}

func NewIntPriorityQueue(prior func(a, b int) bool) *IntPriorityQueue {
	return &IntPriorityQueue{
		queue: NewPriorityQueue(func(a, b interface{}) bool {
			return prior(a.(int), b.(int))
		}),
	}
}

func NewSmallerIntPriorityQueue() *IntPriorityQueue {
	return NewIntPriorityQueue(func(a, b int) bool { return a < b })
}

func NewBiggerIntPriorityQueue() *IntPriorityQueue {
	return NewIntPriorityQueue(func(a, b int) bool { return a > b })
}

// Push は優先度付きキューに要素を一つ追加します。
func (pq IntPriorityQueue) Push(value int) {
	pq.queue.Push(value)
}

// Pop は優先度付きキューから要素を一つ取り出します。
func (pq IntPriorityQueue) Pop() int {
	return pq.queue.Pop().(int)
}

// Top は優先度つきキューの先頭要素を返します。
func (pq IntPriorityQueue) Top() int {
	return pq.queue.Top().(int)
}

// Empty は優先度付きキューが空かどうかを判断します。
func (pq IntPriorityQueue) Empty() bool {
	return pq.queue.Empty()
}

// HasElementsはこの優先度付きキューが要素を持つかどうかを判断します。
func (pq IntPriorityQueue) HasElements() bool {
	return pq.queue.HasElements()
}
