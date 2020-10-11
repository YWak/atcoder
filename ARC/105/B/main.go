package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// N := nextInt()
	N := 1000000
	A := make([]int, N)

	for i := 0; i < N; i++ {
		A[i] = i + 1
		// A[i] = nextInt()
	}
	solve(N, A)
}

func solve(N int, A []int) {
	hist := map[int]bool{}

	lower := pqNew1()
	upper := pqNew2()

	for i := 0; i < N; i++ {
		a := A[i]
		_, exists := hist[a]

		if !exists {
			hist[a] = true
			lower.Push(a)
			upper.Push(a)
		}
	}

	for {
		x := lower.Pop()
		X := upper.Pop()

		if x == X {
			fmt.Println(x)
			return
		}

		X2 := X - x
		lower.Push(x)
		_, exists := hist[int(X2)]

		if exists {
			continue
		}

		upper.Push(X2)

		if x != X2 {
			lower.Push(X2)
		}
	}
}

// PQList1 は 優先度付きキューの本体
type PQList1 []int

// prior は pq[i]の方が優先度が高いかどうかを判断します。
func (pq PQList1) prior(i, j int) bool {
	return pq[i] < pq[j] // 小さいもの優先とする
}

// PriorityQueue1 は優先度付きキューを表す
type PriorityQueue1 struct {
	queue *PQList1
}

func pqNew1() PriorityQueue1 {
	l := make(PQList1, 0, 100)
	return PriorityQueue1{queue: &l}
}

// Push は優先度付きキューに要素を一つ追加します。
func (pq PriorityQueue1) Push(value int) {
	heap.Push(pq.queue, value)
}

// Pop は優先度付きキューから要素を一つ取り出します。
func (pq PriorityQueue1) Pop() int {
	return heap.Pop(pq.queue).(int)
}

// Empty は優先度付きキューが空かどうかを判断します。
func (pq PriorityQueue1) Empty() bool {
	return len(*pq.queue) == 0
}

// Swap は要素を交換します。
func (pq PQList1) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Less は要素を比較し、pq[i] < pq[j]かどうかを判断します
func (pq PQList1) Less(i, j int) bool {
	return pq.prior(i, j)
}

// Len は要素の数を返します。
func (pq PQList1) Len() int {
	return len(pq)
}

// Pop は要素を取り出して返します。
func (pq *PQList1) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Push は要素を追加します。
func (pq *PQList1) Push(item interface{}) {
	*pq = append(*pq, item.(int))
}

// PQList2 は 優先度付きキューの本体
type PQList2 []int

// prior は pq[i]の方が優先度が高いかどうかを判断します。
func (pq PQList2) prior(i, j int) bool {
	return pq[i] > pq[j] // 大きいもの優先とする
}

// PriorityQueue2 は優先度付きキューを表す
type PriorityQueue2 struct {
	queue *PQList2
}

func pqNew2() PriorityQueue2 {
	l := make(PQList2, 0, 100)
	return PriorityQueue2{queue: &l}
}

// Push は優先度付きキューに要素を一つ追加します。
func (pq PriorityQueue2) Push(value int) {
	heap.Push(pq.queue, value)
}

// Pop は優先度付きキューから要素を一つ取り出します。
func (pq PriorityQueue2) Pop() int {
	return heap.Pop(pq.queue).(int)
}

// Empty は優先度付きキューが空かどうかを判断します。
func (pq PriorityQueue2) Empty() bool {
	return len(*pq.queue) == 0
}

// Swap は要素を交換します。
func (pq PQList2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Less は要素を比較し、pq[i] < pq[j]かどうかを判断します
func (pq PQList2) Less(i, j int) bool {
	return pq.prior(i, j)
}

// Len は要素の数を返します。
func (pq PQList2) Len() int {
	return len(pq)
}

// Pop は要素を取り出して返します。
func (pq *PQList2) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Push は要素を追加します。
func (pq *PQList2) Push(item interface{}) {
	*pq = append(*pq, item.(int))
}

var stdin = initStdin()

func initStdin() *bufio.Scanner {
	bufsize := 1 * 1024 * 1024 // 1 MB
	var stdin = bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, bufsize), bufsize)
	stdin.Split(bufio.ScanWords)
	return stdin
}

func nextString() string {
	stdin.Scan()
	return stdin.Text()
}

func nextBytes() []byte {
	stdin.Scan()
	return stdin.Bytes()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func nextInt64() int64 {
	i, _ := strconv.ParseInt(nextString(), 10, 64)
	return i
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
