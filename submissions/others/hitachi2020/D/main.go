package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	N := nextInt()
	T := nextInt64()
	queue := pqNew()

	for i := 0; i < N; i++ {
		a := nextInt64()
		b := nextInt64()

		if a == 0 {
			queue.Push(shop{a: a, b: b, t: T - b})
		} else {
			queue.Push(shop{a: a, b: b, t: (T - b) / a})
		}
	}

	c := 0
	time := int64(0)
	println(len(queue.queue), queue.size)

	for i := 0; i < N; i++ {
		println(i)
		s := queue.Pop()
		t := s.a*time + s.b

		if time+t > T {
			break
		}
		c++
		time += t
	}

	fmt.Println(c)
}

// PriorityQueue は優先度付きキューを表す
type PriorityQueue struct {
	queue []shop
	size  int
}

func pqNew() PriorityQueue {
	return PriorityQueue{queue: make([]shop, 10), size: 0}
}
func (pq PriorityQueue) less(i, j int) bool {
	return pq.queue[i].t < pq.queue[j].t // 小さい方を使用する
}

// Push は優先度付きキューに要素を一つ追加します。
func (pq PriorityQueue) Push(value shop) {
	if len(pq.queue)-1 == pq.size {
		pq.queue = append(pq.queue, value)
	} else {
		pq.queue[pq.size] = value
	}

	i := pq.size
	parent := pq.parent(i)

	for i > 0 && pq.less(parent, i) {
		pq.swap(parent, i)
		i := parent
		parent = pq.parent(i)
	}

	pq.size++
}

// Pop は優先度付きキューから要素を一つ取り出します。
func (pq PriorityQueue) Pop() shop {
	if pq.size == 0 {
		panic("Priority Queue is Empty")
	}
	ret := pq.queue[0]
	pq.size--
	pq.queue[0] = pq.queue[pq.size]

	i := 0
	left := pq.left(i)
	right := pq.right(i)

	for right <= pq.size {
		l := pq.less(i, left)
		r := pq.less(i, right)
		if l && (!r || pq.less(right, left)) {
			pq.swap(i, left)
			i = left
		} else if r {
			pq.swap(i, right)
			i = right
		} else {
			break
		}
		left = pq.left(i)
		right = pq.right(i)
	}

	return ret
}

// Empty は優先度付きキューが空かどうかを判断します。
func (pq PriorityQueue) Empty() bool {
	return pq.size == 0
}

func (pq PriorityQueue) swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}
func (pq PriorityQueue) parent(i int) int {
	return (i - 1) / 2
}
func (pq PriorityQueue) left(i int) int {
	return i*2 + 1
}
func (pq PriorityQueue) right(i int) int {
	return i*2 + 2
}

type shop struct {
	a int64
	b int64
	t int64
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
