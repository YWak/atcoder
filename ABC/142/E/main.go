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

	p := PrimalDual{N: 1 + M + N + 1, graph: make([][]edge, 1+M+N+1)} // S + 鍵 + 宝箱 + T
	for i := 1; i <= M; i++ {
		a := nextInt()
		b := nextInt()

		addEdge(p, 0, i, b, a)

		for j := 0; j < b; j++ {
			c := nextInt()
			addEdge(p, i, M+c, 1, 0)
		}
	}
	for i := 1; i <= N; i++ {
		addEdge(p, M+i, 1+M+N, 1, 0)
	}

	fmt.Println(minCostFlow(p, 0, len(p.graph)-1, inf))
}

var inf = 1 << 29

// PrimalDual is min cost flow graph
type PrimalDual struct {
	N     int
	graph [][]edge
}

type edge struct {
	to   int
	cap  int
	cost int
	rev  int
}

func addEdge(p PrimalDual, from, to, cap, cost int) {
	p.graph[from] = append(p.graph[from], edge{to, cap, cost, len(p.graph[to])})
	p.graph[to] = append(p.graph[to], edge{from, 0, -cost, len(p.graph[from]) - 1})
}

func minCostFlow(p PrimalDual, start, goal, flow int) int {
	prevNode := make([]int, p.N)
	prevEdge := make([]int, p.N)
	potential := make([]int, p.N)
	total := 0

	for flow > 0 {
		dist := make([]int, p.N)
		for i := 0; i < p.N; i++ {
			dist[i] = inf
		}

		pq := make(pqueue, p.N)
		heap.Push(&pq, &node{value: start, priority: 0})
	}

	return total
}

type node struct {
	value    int
	priority int
	index    int
}

type pqueue []*node

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
