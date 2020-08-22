package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var H int
var W int
var Ch int
var Cw int
var Dh int
var Dw int
var S [][]byte
var m [][]int

func main() {
	H = nextInt()
	W = nextInt()
	Ch = nextInt() - 1
	Cw = nextInt() - 1
	Dh = nextInt() - 1
	Dw = nextInt() - 1

	S = make([][]byte, H)
	m = make([][]int, H)

	max := H*W + 1
	for i := 0; i < H; i++ {
		S[i] = nextBytes()
		m[i] = make([]int, W)

		for j := 0; j < W; j++ {
			m[i][j] = max
		}
	}
	m[Ch][Cw] = 0

	queue := pqNew()
	queue.push(P{Ch, Cw, 0})

	dir := []P{
		P{-1, 0, 0},
		P{+1, 0, 0},
		P{0, -1, 0},
		P{0, +1, 0},
	}

	for !queue.Empty() {
		p := queue.pop()
		if p.h == Dh && p.w == Dw {
			fmt.Println(p.c)
			return
		}
		// A type
		for i := 0; i < len(dir); i++ {
			d := dir[i]
			n := P{p.h + d.h, p.w + d.w, p.c}
			if out(n) {
				continue
			}
			if m[n.h][n.w] <= n.c {
				continue
			}
			m[n.h][n.w] = n.c
			queue.push(n)
		}

		// B type
		for h := -2; h <= +2; h++ {
			for w := -2; w <= +2; w++ {
				if h == 0 && w == 0 {
					continue
				}

				n := P{p.h + h, p.w + w, p.c + 1}
				if out(n) {
					continue
				}
				if m[n.h][n.w] <= n.c {
					continue
				}
				m[n.h][n.w] = n.c
				queue.push(n)
			}
		}
	}

	fmt.Println(-1)
}

func out(p P) bool {
	return p.h < 0 || p.w < 0 || p.h >= H || p.w >= W || S[p.h][p.w] == '#'
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

type PQList []P

// PriorityQueue は優先度付きキューを表す
type PriorityQueue struct {
	queue *PQList
	size  int
}

func pqNew() PriorityQueue {
	l := make(PQList, 0, 100)
	pq := PriorityQueue{
		queue: &l,
		size:  0,
	}
	return pq
}

// Push は優先度付きキューに要素を一つ追加します。
func (pq PriorityQueue) push(value P) {
	heap.Push(pq.queue, value)
}

// Pop は優先度付きキューから要素を一つ取り出します。
func (pq PriorityQueue) pop() P {
	return heap.Pop(pq.queue).(P)
}

// Empty は優先度付きキューが空かどうかを判断します。
func (pq PriorityQueue) Empty() bool {
	return len(*pq.queue) == 0
}

func (pq PQList) Len() int {
	return len(pq)
}

func (pq PQList) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQList) Less(i, j int) bool {
	return pq.prior(i, j)
}

func (pq *PQList) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func (pq *PQList) Push(item interface{}) {
	*pq = append(*pq, item.(P))
}

func (pq PQList) prior(i, j int) bool {
	pi := pq[i]
	pj := pq[j]

	if pi.c != pj.c {
		return pi.c < pj.c
	}
	di := abs(pi.h-Dh) + abs(pi.w-Dw)
	dj := abs(pj.h-Dh) + abs(pj.w-Dw)
	return di < dj
}

type P struct {
	h int
	w int
	c int
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
	fmt.Fprintln(os.Stderr, args)
}
