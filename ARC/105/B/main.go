package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	hist := map[int]bool{}

	lower := pqNew(-1)
	upper := pqNew(1)

	for i := 0; i < N; i++ {
		a := nextInt()

		_, exists := hist[a]

		if !exists {
			hist[a] = true
			lower.Push(PQItem(a))
			upper.Push(PQItem(a))
		}
	}

	for {
		x := lower.Pop()
		X := upper.Pop()

		if x == X {
			fmt.Println(x)
			return
		}

		lower.Push(x)
		lower.Push(X - x)
		upper.Push(X - x)
	}
}

// PQItem は優先度付きキューに保存される要素
type PQItem int

// PQList は 優先度付きキューの本体
type PQList struct {
	list  []PQItem
	order int
}

// prior は pq[i]の方が優先度が高いかどうかを判断します。
func (pq PQList) prior(i, j int) bool {
	return int(pq.list[i])*pq.order > int(pq.list[j])*pq.order // 大きいもの優先とする
}

// PriorityQueue は優先度付きキューを表す
type PriorityQueue struct {
	queue *PQList
}

func pqNew(order int) PriorityQueue {
	l := PQList{make([]PQItem, 0, 100), order}
	return PriorityQueue{queue: &l}
}

// Push は優先度付きキューに要素を一つ追加します。
func (pq PriorityQueue) Push(value PQItem) {
	heap.Push(pq.queue, value)
}

// Pop は優先度付きキューから要素を一つ取り出します。
func (pq PriorityQueue) Pop() PQItem {
	return heap.Pop(pq.queue).(PQItem)
}

// Empty は優先度付きキューが空かどうかを判断します。
func (pq PriorityQueue) Empty() bool {
	return len(pq.queue.list) == 0
}

// Swap は要素を交換します。
func (pq PQList) Swap(i, j int) {
	pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
}

// Less は要素を比較し、pq[i] < pq[j]かどうかを判断します
func (pq PQList) Less(i, j int) bool {
	return pq.prior(i, j)
}

// Len は要素の数を返します。
func (pq PQList) Len() int {
	return len(pq.list)
}

// Pop は要素を取り出して返します。
func (pq *PQList) Pop() interface{} {
	old := pq.list
	n := len(old)
	item := old[n-1]
	pq.list = old[:n-1]
	return item
}

// Push は要素を追加します。
func (pq *PQList) Push(item interface{}) {
	pq.list = append(pq.list, item.(PQItem))
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
