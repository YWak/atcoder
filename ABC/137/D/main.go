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
	AB := map[int][]int{}

	for i := 0; i < N; i++ {
		a := nextInt()
		b := nextInt()

		AB[a] = append(AB[a], b)
	}
	pq := make(pqueue, 0)

	c := 0
	for i := 0; i < M; i++ {
		// (i+1)日後に報酬を受け取れるアルバイトを追加する
		for j := 0; j < len(AB[i+1]); j++ {
			b := AB[i+1][j]
			heap.Push(&pq, &item{value: b, priority: b})
		}
		if len(pq) != 0 {
			v := heap.Pop(&pq)
			c += (*v.(*item)).value
		}
	}

	fmt.Println(c)
}

type item struct {
	value    int
	priority int
	index    int
}

type pqueue []*item

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
	i := x.(*item)
	i.index = n
	*pq = append(*pq, x.(*item))
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
