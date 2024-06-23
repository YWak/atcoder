package ds

// Stackは先入れ後出しのデータ構造を表現します。
type Stack[T any] struct {
	arr []T

	// PushはこのStackの末尾に要素を追加します。
	Push func(value T)

	// PopはこのStackの末尾から要素を取り出し、その要素を返します。
	Pop func() T

	// PeekはこのStackの末尾の要素を返します。
	// 末尾の要素は削除されません。
	Peek func() T

	// CountはこのStackに保存されている要素の数を返します。
	Count func() int

	// HasElementsはこのStackに要素が設定されているかを判断します。
	HasElements func() bool

	// IsEmptyはこのStackが空かどうかを判断します。
	IsEmpty func() bool
}

func NewStack[T any]() *Stack[T] {
	stack := Stack[T]{}
	stack.Push = func(value T) {
		stack.arr = append(stack.arr, value)
	}
	stack.Pop = func() T {
		if len(stack.arr) == 0 {
			panic("Pop() is called for empty stack")
		}

		value := stack.Peek()
		stack.arr = stack.arr[:len(stack.arr)-1]
		return value
	}
	stack.Peek = func() T {
		if len(stack.arr) == 0 {
			panic("Peek() is called for empty stack")
		}

		return stack.arr[len(stack.arr)-1]
	}
	stack.Count = func() int {
		return len(stack.arr)
	}
	stack.HasElements = func() bool {
		return stack.Count() > 0
	}
	stack.IsEmpty = func() bool {
		return stack.Count() == 0
	}
	return &stack
}
