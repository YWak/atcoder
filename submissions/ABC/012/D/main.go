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

// INF は最大値を表す数
const INF = int(1e15)

func main() {
	N := nextInt()
	M := nextInt()
	g := NewGraph(N)

	for i := 0; i < M; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		t := nextInt()
		g.AddWeightedEdge(a, b, t)
		g.AddWeightedEdge(b, a, t)
	}
	dist := g.WarshallFloyd()

	ans := INF
	for i := 0; i < N; i++ {
		m := 0
		for j := 0; j < N; j++ {
			m = max(m, dist[i][j])
		}
		ans = min(ans, m)
	}
	// debug()
	// for i := 0; i < N; i++ {
	// 	for j := i; j < N; j++ {
	// 		debug(i+1, "-", j+1, "=", dist[i][j])
	// 	}
	// }
	fmt.Println(ans)
}

// Graph はグラフを表現する構造です
type Graph struct {
	// 隣接リスト
	list [][]Edge
}

// Edge は辺を表現する構造体です
type Edge struct {
	to     int
	weight int
}

// NewGraph はグラフを作成します
func NewGraph(n int) *Graph {
	return &(Graph{make([][]Edge, n)})
}

// AddEdge は辺を追加します
func (g *Graph) AddEdge(s, t int) {
	g.AddWeightedEdge(s, t, 1)
}

// AddWeightedEdge は重み付きの辺を追加します。
func (g *Graph) AddWeightedEdge(s, t, w int) {
	g.list[s] = append(g.list[s], Edge{t, w})
}

// DijkstraNode は ダイクストラ法を使用するときに使うノード
type DijkstraNode struct {
	node int
	cost int
}

// DijkstraPriorityQueue はダイクストラ法を使用するときに使う優先度付きキュー
type DijkstraPriorityQueue []*DijkstraNode

func (pq DijkstraPriorityQueue) Len() int           { return len(pq) }
func (pq DijkstraPriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq DijkstraPriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

// Push はpqに要素を追加する
func (pq *DijkstraPriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*DijkstraNode)) }

// Pop はpqから要素を取得する
func (pq *DijkstraPriorityQueue) Pop() interface{} {
	o := *pq
	n := len(o) - 1
	item := o[n]
	*pq = o[0:n]
	return item
}

// Dijkstra はsからtへの最短距離と最短ルートを返します。
// 重みが負の辺があるときには使用できません。
// 計算量: |V| + |E|log|V|
func (g *Graph) Dijkstra(s, t int) (int, []int) {
	n := len(g.list)
	pq := make(DijkstraPriorityQueue, 0)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		var c int
		if i == s {
			c = 0
		} else {
			c = INF
		}
		cost[i] = c
		heap.Push(&pq, &DijkstraNode{i, c})
	}

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*DijkstraNode)
		for i := 0; i < len(g.list[u.node]); i++ {
			v := g.list[u.node][i]
			c := cost[u.node] + v.weight
			if cost[v.to] > c {
				cost[v.to] = c
				heap.Push(&pq, &DijkstraNode{v.to, c})
			}
		}
	}

	return cost[t], []int{}
}

// BellmanFord はsからtへの最短ルートを返します。
func (g *Graph) BellmanFord(s, t int) {
}

// WarshallFloyd は全点対の最短ルートを返します。
func (g *Graph) WarshallFloyd() [][]int {
	n := len(g.list)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = INF
			}
		}
		for j := 0; j < len(g.list[i]); j++ {
			k := g.list[i][j]
			d[i][k.to] = k.weight
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	return d
}

// Dinic はsからtへの最小費用流を返します。

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
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

// toi は byteの数値をintに変換します。
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

func nextFloat() float64 {
	f, _ := strconv.ParseFloat(nextString(), 64)
	return f
}

// ==================================================
// 数値操作
// ==================================================

// max は aとbのうち大きい方を返します。
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min は aとbのうち小さい方を返します。
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// abs は aの絶対値を返します。
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// pow は aのb乗を返します。
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

// toLowerCase は sをすべて小文字にした文字列を返します。
func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// toUpperCase は sをすべて大文字にした文字列を返します。
func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

// isLower はbが小文字かどうかを判定します
func isLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}

// isUpper はbが大文字かどうかを判定します
func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

// ==================================================
// 構造体
// ==================================================

// Point は 座標を表す構造体です。
type Point struct {
	x int
	y int
}

// Pointf は座標を表す構造体です。
type Pointf struct {
	x float64
	y float64
}
