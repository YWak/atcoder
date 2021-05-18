//lint:file-ignore U1000 using template
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)

// INF18 は最大値を表す数
const INF18 = int(1e18)

// INF9 は最大値を表す数
const INF9 = int(1e9)

var in *In
var out *Out

func calc() {
	h, w := in.NextInt2()
	rs, cs := in.NextInt2()
	rt, ct := in.NextInt2()
	rs--
	cs--
	rt--
	ct--

	s := make([]string, h)
	ew := make([][]int, h)
	eh := make([][]int, h)
	for i := 0; i < h; i++ {
		s[i] = in.NextString()
		ew[i] = make([]int, w)
		eh[i] = make([]int, w)
	}
	// エッジの検出
	c := 0
	for i := 0; i < h; i++ {
		inLine := false
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				inLine = false
			} else {
				if !inLine {
					c++
				}
				inLine = true
			}
			if inLine {
				ew[i][j] = c
			}
		}
		c++
	}
	for j := 0; j < w; j++ {
		inLine := false
		for i := 0; i < h; i++ {
			if s[i][j] == '#' {
				inLine = false
			} else {
				if !inLine {
					c++
				}
				inLine = true
			}
			if inLine {
				eh[i][j] = c
			}
		}
		c++
	}

	graph := NewGraph(c + 2)
	start := c
	goal := c + 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				continue
			}
			a, b := eh[i][j], ew[i][j]
			graph.AddWeightedEdge(a, b, 1)
			graph.AddWeightedEdge(b, a, 1)
		}
	}
	graph.AddWeightedEdge(start, eh[rs][cs], 0)
	graph.AddWeightedEdge(start, ew[rs][cs], 0)
	graph.AddWeightedEdge(eh[rt][ct], goal, 0)
	graph.AddWeightedEdge(ew[rt][ct], goal, 0)
	fmt.Println(graph.Dijkstra(start, goal))
}

func main() {
	in, out = InitIo()
	defer out.writer.Flush()

	calc()
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
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

// Dijkstra はsからtへの最短距離を返します。
// 重みが負の辺があるときには使用できません。
// 計算量: |V| + |E|log|V|
func (g *Graph) Dijkstra(s, t int) int {
	n := len(g.list)
	pq := make(DijkstraPriorityQueue, 0)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		var c int
		if i == s {
			c = 0
		} else {
			c = INF18
		}
		cost[i] = c
		heap.Push(&pq, &DijkstraNode{i, c})
	}

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*DijkstraNode)
		if u.node == t {
			break
		}
		for i := 0; i < len(g.list[u.node]); i++ {
			v := g.list[u.node][i]
			c := cost[u.node] + v.weight
			if cost[v.to] > c {
				cost[v.to] = c
				heap.Push(&pq, &DijkstraNode{v.to, c})
			}
		}
	}

	return cost[t]
}

// DijkstraAll はsから全点への最短距離を返します。
// 重みが負の辺があるときには使用できません。
// 計算量: |V| + |E|log|V|
func (g *Graph) DijkstraAll(s int) []int {
	n := len(g.list)
	pq := make(DijkstraPriorityQueue, 0)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		var c int
		if i == s {
			c = 0
		} else {
			c = INF18
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

	return cost
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
				d[i][j] = INF18
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

// ==================================================
// 入出力操作
// ==================================================
type In struct {
	reader *bufio.Scanner
}
type Out struct {
	writer *bufio.Writer
}

func InitIo() (*In, *Out) {
	bufsize := 4 * 1024 * 1024 // 1 MB
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, bufsize), bufsize)
	in.Split(bufio.ScanWords)

	out := bufio.NewWriterSize(os.Stdout, bufsize)

	return &In{in}, &Out{out}
}

func (in *In) NextString() string {
	in.reader.Scan()
	return in.reader.Text()
}

// 遅いから極力使わない。
func (in *In) NextBytes() []byte {
	return []byte(in.NextString())
}

func (in *In) NextInt() int {
	i, _ := strconv.Atoi(in.NextString())
	return i
}

func (in *In) NextInt2() (int, int) {
	return in.NextInt(), in.NextInt()
}

func (in *In) NextInt3() (int, int, int) {
	return in.NextInt(), in.NextInt(), in.NextInt()
}

func (in *In) nextInt4() (int, int, int, int) {
	return in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
}

func (in *In) NextInts(n int) sort.IntSlice {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = in.NextInt()
	}
	return sort.IntSlice(a)
}

// toi は byteの数値をintに変換します。
func toi(b byte) int {
	return int(b - '0')
}

func (in *In) NextLongIntAsArray() []int {
	s := in.NextString()
	l := len(s)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = toi(s[i])
	}

	return arr
}

func (in *In) NextFloat() float64 {
	f, _ := strconv.ParseFloat(in.NextString(), 64)
	return f
}

// NextFloatAsInt は 数を 10^base 倍した整数値を取得します。
func (in *In) NextFloatAsInt(base int) int {
	s := in.NextString()
	index := strings.IndexByte(s, '.')
	if index == -1 {
		n, _ := strconv.Atoi(s)
		return n * pow(10, base)
	}
	for s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	s1, s2 := s[:index], s[index+1:]
	n, _ := strconv.Atoi(s1)
	m, _ := strconv.Atoi(s2)
	return n*pow(10, base) + m*pow(10, base-len(s2))
}

func (out *Out) Println(a ...interface{}) {
	fmt.Fprintln(out.writer, a...)
}

func (out *Out) Printf(format string, a ...interface{}) {
	fmt.Fprintf(out.writer, format, a...)
}

func (out *Out) PrintStringsln(a []string) {
	b := make([]interface{}, len(a))
	for i, v := range a {
		b[i] = v
	}
	out.Println(b...)
}

func (out *Out) PrintIntsLn(a []int) {
	b := make([]interface{}, len(a))
	for i, v := range a {
		b[i] = v
	}
	out.Println(b...)
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

// divceil は a/b の結果を正の無限大に近づけるように丸めて返します。
func divceil(a, b int) int {
	if a%b == 0 || a/b < 0 {
		return a / b
	}
	return (a + b - 1) / b
}

// divfloor は a/b の結果を負の無限大に近づけるように丸めて返します。
func divfloor(a, b int) int {
	if a%b == 0 || a/b > 0 {
		return a / b
	}
	if b < 0 {
		a, b = -a, -b
	}
	return (a - b + 1) / b
}

// powmod は (x^n) mod m を返します。
func powmod(x, n, m int) int {
	ans := 1
	for n > 0 {
		if n%2 == 1 {
			ans = (ans * x) % m
		}
		x = (x * x) % m
		n /= 2
	}
	return ans
}

// ch は condがtrueのときok, falseのときngを返します。
func ch(cond bool, ok, ng int) int {
	if cond {
		return ok
	}
	return ng
}

func mul(a, b int) (int, int) {
	if a < 0 {
		a, b = -a, -b
	}
	if a == 0 || b == 0 {
		return 0, 0
	} else if a > 0 && b > 0 && a > math.MaxInt64/b {
		return 0, +1
	} else if a > math.MinInt64/b {
		return 0, -1
	}
	return a * b, 0
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

func xor(a, b bool) bool { return a != b }

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
// 配列
// ==================================================0
func reverse(arr *[]interface{}) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func reverseInt(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func uniqueInt(arr []int) []int {
	hist := map[int]bool{}
	j := 0
	for i := 0; i < len(arr); i++ {
		if hist[arr[i]] {
			continue
		}

		a := arr[i]
		arr[j] = a
		hist[a] = true
		j++
	}
	return arr[:j]
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
