package ds

// Stackは先入れ後出しのデータ構造を表現します。
type Stack[T any] struct {
	arr []T
}

// PushはこのStackの末尾に要素を追加します。
func (stack *Stack[T]) Push(value T) {
	stack.arr = append(stack.arr, value)
}

// PopはこのStackの末尾から要素を取り出し、その要素を返します。
func (stack *Stack[T]) Pop() T {
	if len(stack.arr) == 0 {
		panic("Pop() is called for empty stack")
	}

	value := stack.Peek()
	stack.arr = stack.arr[:len(stack.arr)-1]
	return value
}

// PeekはこのStackの末尾の要素を返します。
// 末尾の要素は削除されません。
func (stack *Stack[T]) Peek() T {
	if len(stack.arr) == 0 {
		panic("Peek() is called for empty stack")
	}

	return stack.arr[len(stack.arr)-1]
}

// CountはこのStackに保存されている要素の数を返します。
func (stack *Stack[T]) Count() int {
	return len(stack.arr)
}

// HasElementsはこのStackに要素が設定されているかを判断します。
func (stack *Stack[T]) HasElements() bool {
	return stack.Count() > 0
}

// IsEmptyはこのStackが空かどうかを判断します。
func (stack *Stack[T]) IsEmpty() bool {
	return !stack.HasElements()
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}
