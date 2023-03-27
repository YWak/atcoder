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
	M := nextInt()
	pq := make(pqueue, N)

	for i := 0; i < N; i++ {
		a := nextInt64()
		item := cost{value: a, priority: a}
		pq[i] = &item
	}
	heap.Init(&pq)

	for i := 0; i < M; i++ {
		item := heap.Pop(&pq).(*cost)
		item.value /= 2
		item.priority /= 2
		heap.Push(&pq, item)
	}
	all := int64(0)
	for i := 0; i < N; i++ {
		item := heap.Pop(&pq)
		all += (item.(*cost)).value
	}

	fmt.Println(all)
}

type cost struct {
	value    int64
	priority int64
	index    int
}

type pqueue []*cost

// キューの長さ
func (pq pqueue) Len() int { return len(pq) }

// キューの大小比較
// 基本は大きい方を優先する
func (pq pqueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }

// キューの内容交換
func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// キューへの追加
func (pq *pqueue) Push(x interface{}) {
	n := len(*pq)
	i := x.(*cost)
	i.index = n
	*pq = append(*pq, x.(*cost))
}

// キューからの取り出し
func (pq *pqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	if n == 0 {
		fmt.Println("queue is empty")
		return nil
	}
	i := old[n-1]
	i.index = -1
	*pq = old[0 : n-1]
	return i
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
