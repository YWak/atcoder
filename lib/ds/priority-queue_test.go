package ds_test

import (
	"sort"
	"testing"

	"github.com/ywak/atcoder/lib/ds"
)

func TestPriorityQueueBasicOperations(t *testing.T) {
	// 最小ヒープのテスト (小さい値が優先)
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	// 初期状態のテスト
	if !pq.IsEmpty() {
		t.Error("新しいキューは空であるべき")
	}
	if pq.HasElements() {
		t.Error("新しいキューは要素を持たないべき")
	}
	if pq.Count() != 0 {
		t.Error("新しいキューの要素数は0であるべき")
	}

	// 要素の追加
	values := []int{5, 2, 8, 1, 9, 3}
	for _, v := range values {
		pq.Push(v)
	}

	// 状態確認
	if pq.IsEmpty() {
		t.Error("要素を追加した後は空でないべき")
	}
	if !pq.HasElements() {
		t.Error("要素を追加した後は要素を持つべき")
	}
	if pq.Count() != len(values) {
		t.Errorf("要素数が間違っている: expected %d, got %d", len(values), pq.Count())
	}

	// Peekのテスト（最小値が返されるべき）
	minValue := pq.Peek()
	if minValue != 1 {
		t.Errorf("Peekが最小値を返していない: expected 1, got %d", minValue)
	}

	// 要素数が変わらないことを確認
	if pq.Count() != len(values) {
		t.Error("Peekで要素数が変わってはいけない")
	}
}

func TestPriorityQueueOrdering(t *testing.T) {
	// 最小ヒープのテスト
	t.Run("MinHeap", func(t *testing.T) {
		pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

		values := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
		for _, v := range values {
			pq.Push(v)
		}

		expected := make([]int, len(values))
		copy(expected, values)
		sort.Ints(expected) // 昇順ソート

		result := []int{}
		for !pq.IsEmpty() {
			result = append(result, pq.Pop())
		}

		if len(result) != len(expected) {
			t.Errorf("結果の長さが異なる: expected %d, got %d", len(expected), len(result))
		}

		for i, v := range result {
			if v != expected[i] {
				t.Errorf("順序が間違っている: index %d, expected %d, got %d", i, expected[i], v)
			}
		}
	})

	// 最大ヒープのテスト
	t.Run("MaxHeap", func(t *testing.T) {
		pq := ds.NewPriorityQueue(func(a, b int) bool { return a > b })

		values := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
		for _, v := range values {
			pq.Push(v)
		}

		expected := make([]int, len(values))
		copy(expected, values)
		sort.Sort(sort.Reverse(sort.IntSlice(expected))) // 降順ソート

		result := []int{}
		for !pq.HasElements() == false {
			result = append(result, pq.Pop())
		}

		if len(result) != len(expected) {
			t.Errorf("結果の長さが異なる: expected %d, got %d", len(expected), len(result))
		}

		for i, v := range result {
			if v != expected[i] {
				t.Errorf("順序が間違っている: index %d, expected %d, got %d", i, expected[i], v)
			}
		}
	})
}

func TestPriorityQueueSingleElement(t *testing.T) {
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	pq.Push(42)

	if pq.Count() != 1 {
		t.Errorf("要素数が間違っている: expected 1, got %d", pq.Count())
	}

	if pq.Peek() != 42 {
		t.Errorf("Peekの結果が間違っている: expected 42, got %d", pq.Peek())
	}

	if pq.Pop() != 42 {
		t.Errorf("Popの結果が間違っている: expected 42, got %d", pq.Pop())
	}

	if !pq.IsEmpty() {
		t.Error("要素を取り出した後は空であるべき")
	}
}

func TestPriorityQueueDuplicateValues(t *testing.T) {
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	// 重複値をプッシュ
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	for _, v := range values {
		pq.Push(v)
	}

	expected := make([]int, len(values))
	copy(expected, values)
	sort.Ints(expected)

	result := []int{}
	for pq.HasElements() {
		result = append(result, pq.Pop())
	}

	if len(result) != len(expected) {
		t.Errorf("結果の長さが異なる: expected %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("順序が間違っている: index %d, expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestPriorityQueueHeapProperty(t *testing.T) {
	// ヒープの性質を確認するためのテスト
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	// ランダムな値を追加
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, v := range values {
		pq.Push(v)
	}

	// 各Popで最小値が取得されることを確認
	lastPop := -1
	for pq.HasElements() {
		current := pq.Pop()
		if current < lastPop {
			t.Errorf("ヒープの性質が破れている: 前回のPop %d > 現在のPop %d", lastPop, current)
		}
		lastPop = current
	}
}

func TestPriorityQueueStressTest(t *testing.T) {
	// 大きなデータセットでのテスト
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	// 1000個の要素をプッシュ
	n := 1000
	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = n - i // 逆順でプッシュ
		pq.Push(values[i])
	}

	// すべてポップして順序を確認
	for i := 1; i <= n; i++ {
		if pq.Count() != n-i+1 {
			t.Errorf("要素数が間違っている: expected %d, got %d", n-i+1, pq.Count())
		}

		popped := pq.Pop()
		if popped != i {
			t.Errorf("順序が間違っている: expected %d, got %d", i, popped)
		}
	}

	if !pq.IsEmpty() {
		t.Error("すべてポップした後は空であるべき")
	}
}

// エラーケースのテスト（パニックが発生する可能性がある操作）
func TestPriorityQueueErrorCases(t *testing.T) {
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	// 空のキューでPeekを試行
	t.Run("EmptyPeek", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("空のキューでPeekした際にパニック: %v", r)
			}
		}()

		if pq.IsEmpty() {
			// 空のキューでPeekを呼び出す
			_ = pq.Peek()
			t.Log("空のキューでPeekしても問題なし（またはデフォルト値が返された）")
		}
	})

	// 空のキューでPopを試行
	t.Run("EmptyPop", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("空のキューでPopした際にパニック: %v", r)
			}
		}()

		if pq.IsEmpty() {
			// 空のキューでPopを呼び出す
			_ = pq.Pop()
			t.Log("空のキューでPopしても問題なし（またはデフォルト値が返された）")
		}
	})
}

// デバッグ用のヘルパー関数
func printHeapStructure(t *testing.T, name string, values []int) {
	t.Logf("%s: %v", name, values)
}

// ベンチマークテスト
func BenchmarkPriorityQueuePush(b *testing.B) {
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}
}

func BenchmarkPriorityQueuePop(b *testing.B) {
	pq := ds.NewPriorityQueue(func(a, b int) bool { return a < b })

	// 事前にデータを準備
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}
