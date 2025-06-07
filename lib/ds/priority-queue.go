package ds

// PriorityQueue は優先度付きキューを表す
type PriorityQueue[T any] struct {
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
	list := []T{}

	return &PriorityQueue[T]{
		Push: func(value T) {
			list = append(list, value)
			for leaf := len(list) - 1; leaf > 0; /* nop */ {
				// 親と比べて小さければ値を入れ替えて、処理を継続する
				parent := (leaf - 1) / 2
				if prior(list[parent], list[leaf]) {
					break
				}
				list[leaf], list[parent] = list[parent], list[leaf]
				leaf = parent
			}
		},
		Pop: func() T {
			val := list[0]
			last := len(list) - 1
			list[0] = list[last]
			list = list[:last]

			var next int
			for parent := 0; parent < len(list); parent = next {
				left := parent*2 + 1
				right := parent*2 + 2

				// 子が存在しないので終了
				if left >= len(list) {
					break
				}

				// どちらが大きいか？
				if right >= len(list) || prior(list[left], list[right]) {
					// 左を採用する
					next = left
				} else {
					// 右を採用する
					next = right
				}

				// 現在の値が子より優先度が高いなら終了
				if prior(list[parent], list[next]) {
					break
				}
				list[parent], list[next] = list[next], list[parent]
			}

			return val
		},
		Peek: func() T {
			return list[0]
		},
		IsEmpty: func() bool {
			return len(list) == 0
		},
		HasElements: func() bool {
			return len(list) > 0
		},
		Count: func() int {
			return len(list)
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
