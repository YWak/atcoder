package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

var H int
var W int
var A []string
var c [][]int

func main() {
	H = nextInt()
	W = nextInt()

	A = make([]string, H)
	c = make([][]int, H)
	var s point
	var g point
	telepo := map[byte][]point{}
	MAX := H*W + 1

	for i := 0; i < H; i++ {
		A[i] = nextString()
		c[i] = make([]int, W)

		for j := 0; j < W; j++ {
			c[i][j] = MAX
			if A[i][j] == 'S' {
				s = point{i, j}
			} else if A[i][j] == 'G' {
				g = point{i, j}
			} else if A[i][j] == '.' || A[i][j] == '#' {
				// NOP
			} else {
				telepo[A[i][j]] = append(telepo[A[i][j]], point{i, j})
			}
		}
	}

	queue := pqNew()
	queue.Push(&s)
	c[s.x][s.y] = 0
	dir := []point{
		{-1, +0},
		{+1, +0},
		{+0, -1},
		{+0, +1},
	}

	for !queue.Empty() {
		p := queue.Pop()
		cost := c[p.x][p.y]

		for i := 0; i < len(dir); i++ {
			x := p.x + dir[i].x
			y := p.y + dir[i].y

			if out(x, y) {
				continue
			}
			if c[x][y] <= cost+1 {
				continue
			}
			c[x][y] = cost + 1
			queue.Push(&point{x, y})
		}
		l, exists := telepo[A[p.x][p.y]]
		if exists {
			for i := 0; i < len(l); i++ {
				n := l[i]
				if c[n.x][n.y] > cost+1 {
					c[n.x][n.y] = cost + 1
					queue.Push(&point{n.x, n.y})
				}
			}
		}
	}
	ans := c[g.x][g.y]
	if ans == MAX {
		ans = -1
	}
	fmt.Println(ans)
}
func out(x, y int) bool {
	return !(x >= 0 && x < H && y >= 0 && y < W && A[x][y] != '#')
}

// PQItem は優先度付きキューに保存される要素
type PQItem *point

// PQList は 優先度付きキューの本体
type PQList []PQItem

// prior は pq[i]の方が優先度が高いかどうかを判断します。
func (pq PQList) prior(i, j int) bool {
	p1 := pq[i]
	p2 := pq[j]
	return c[p1.x][p1.y] < c[p2.x][p2.y] // 大きいもの優先とする
}

// PriorityQueue は優先度付きキューを表す
type PriorityQueue struct {
	queue *PQList
}

func pqNew() PriorityQueue {
	l := make(PQList, 0, 100)
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
	return len(*pq.queue) == 0
}

// Swap は要素を交換します。
func (pq PQList) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Less は要素を比較し、pq[i] < pq[j]かどうかを判断します
func (pq PQList) Less(i, j int) bool {
	return pq.prior(i, j)
}

// Len は要素の数を返します。
func (pq PQList) Len() int {
	return len(pq)
}

// Pop は要素を取り出して返します。
func (pq *PQList) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Push は要素を追加します。
func (pq *PQList) Push(item interface{}) {
	*pq = append(*pq, item.(PQItem))
}

// ==================================================
// 入力操作
// ==================================================
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

// 遅いから極力使わない。
func nextBytes() []byte {
	return []byte(nextString())
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func nextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return a
}

func toi(b byte) int {
	return int(b - '0')
}

func nextLongIntAsArray() []int {
	s := nextString()
	l := len(s)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = toi(s[i])
	}

	return arr
}

// ==================================================
// 数値操作
// ==================================================
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// binarysearch は judgeがtrueを返す最小の数値を返します。
func binarysearch(ok, ng int, judge func(int) bool) int {
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if judge(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// ==================================================
// ビット操作
// ==================================================

// nthbit はaのn番目のビットを返します。
func nthbit(a int, n int) int { return int((a >> uint(n)) & 1) }

// popcount はaのうち立っているビットを数えて返します。
func popcount(a int) int {
	return bits.OnesCount(uint(a))
}

// ==================================================
// 文字列操作
// ==================================================
func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

// ==================================================
// 構造体
// ==================================================
type point struct {
	x int
	y int
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
