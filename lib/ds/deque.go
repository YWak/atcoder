package ds

// DequeはStackを2つ使用したDequeの実装です。
type Deque[T any] struct {
	heads *Stack[T]
	tails *Stack[T]
}

type balanceOrder int

const (
	Head balanceOrder = iota
	Tail
)

// PushはStack.Pushの実装
func (queue *Deque[T]) Push(value T) {
	queue.tails.Push(value)
}

// PopはStack.Popの実装
func (queue *Deque[T]) Pop() T {
	queue.balance(Tail)
	return queue.tails.Pop()
}

// PeekはStack.Peekの実装
func (queue *Deque[T]) Peek() T {
	queue.balance(Tail)
	return queue.tails.Peek()
}

// AddはQueue.Addの実装
func (queue *Deque[T]) Add(value T) {
	queue.tails.Push(value)
}

// RemoveはQueue.Removeの実装
func (queue *Deque[T]) Remove() T {
	queue.balance(Head)

	if queue.heads.IsEmpty() {
		panic("Remove() is called for empty Queue")
	}

	return queue.heads.Pop()
}

// GetはQueue.Getの実装
func (queue *Deque[T]) Get() T {
	queue.balance(Head)
	if queue.heads.IsEmpty() {
		panic("Get() is called for empty Queue")
	}

	return queue.heads.Peek()
}

// balanceはheadsとtailsに保存されている要素のバランスを取ります。
func (queue *Deque[T]) balance(order balanceOrder) {
	var src *Stack[T]
	var dst *Stack[T]
	if order == Head {
		// tails -> heads
		src, dst = queue.tails, queue.heads
	} else {
		// heads -> tails
		src, dst = queue.heads, queue.tails
	}

	if dst.HasElements() {
		return
	}

	size := (src.Count() + 1) / 2
	dst.arr = src.arr[:size]
	src.arr = src.arr[size:]
}

// CountはQueue.Countの実装
func (queue *Deque[T]) Count() int {
	return queue.heads.Count() + queue.tails.Count()
}

// HasElementsはQueue.HasElementsの実装
func (queue *Deque[T]) HasElements() bool {
	return queue.heads.HasElements() || queue.tails.HasElements()
}

// IsEmptyはQueue.IsEmptyの実装
func (queue *Deque[T]) IsEmpty() bool {
	return !queue.HasElements()
}

func (queue *Deque[T]) AddFirst(value T) {
	queue.Add(value)
}

func (queue *Deque[T]) AddLast(value T) {
	queue.Push(value)
}

func (queue *Deque[T]) RemoveFirst() T {
	return queue.Remove()
}

func (queue *Deque[T]) RemoveLast() T {
	return queue.Pop()
}

func (queue *Deque[T]) First() T {
	return queue.Get()
}

func (queue *Deque[T]) Last() T {
	return queue.Peek()
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		heads: &Stack[T]{},
		tails: &Stack[T]{},
	}
}
