package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

// INF は最大値を表す数
const INF = int(1e9)

var graph [][]int

func main() {
	N := nextInt()
	M := nextInt()
	graph = make([][]int, N)
	// AB
	for i := 0; i < M; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	K := nextInt()
	C := make([]int, K)
	dist := make([][]int, K)
	for i := 0; i < K; i++ {
		C[i] = nextInt() - 1
	}
	for i := 0; i < K; i++ {
		dist[i] = bfs(C[i], C)
		for j := 0; j < K; j++ {
			if dist[i][j] == -1 {
				fmt.Println(-1)
				return
			}
		}
	}

	dp := make([][]int, K)
	for i := 0; i < K; i++ {
		dp[i] = make([]int, 1<<K)
		for j := 0; j < (1 << K); j++ {
			dp[i][j] = INF
		}
		dp[i][1<<i] = 1
	}
	for c := 1; c < (1 << K); c++ {
		for i := 0; i < K; i++ {
			if (c & (1 << i)) == 0 {
				continue
			}
			bit := c ^ (1 << i)
			for j := 0; j < K; j++ {
				if (bit & (1 << j)) == 0 {
					continue
				}
				dp[i][c] = min(dp[i][c], dp[j][bit]+dist[i][j])
			}
		}
	}
	ans := INF
	for i := 0; i < K; i++ {
		ans = min(ans, dp[i][(1<<K)-1])
	}
	fmt.Println(ans)
}

type pair struct{ a, b int }

func bfs(s int, c []int) []int {
	used := map[int]bool{s: true}
	queue := []pair{pair{s, 0}}
	dist := make([]int, len(graph))
	for i := 0; i < len(dist); i++ {
		dist[i] = -1
	}
	dist[s] = 0

	for len(queue) > 0 {
		p := queue[0]
		d := p.b + 1
		queue = queue[1:]

		for i := 0; i < len(graph[p.a]); i++ {
			j := graph[p.a][i]
			_, ok := used[j]
			if ok {
				continue
			}
			dist[j] = d
			queue = append(queue, pair{j, d})
			used[j] = true
		}
	}
	ret := make([]int, len(c))
	for i := 0; i < len(c); i++ {
		ret[i] = dist[c[i]]
	}
	return ret
}

// UnionFind : UnionFind構造を保持する構造体
type UnionFind struct {
	par  []int // i番目のノードに対応する親
	rank []int // i番目のノードの階層
}

// [0, n)のノードを持つUnion-Findを作る
func ufNew(n int) UnionFind {
	uf := UnionFind{par: make([]int, n), rank: make([]int, n)}

	for i := 0; i < n; i++ {
		uf.par[i] = i
	}

	return uf
}

// Root はxのルートを得る
func (uf *UnionFind) Root(x int) int {
	p := x
	for p != uf.par[p] {
		p = uf.par[p]
	}
	uf.par[x] = p
	return p
}

// Unite はxとyを併合する。集合の構造が変更された(== 呼び出し前は異なる集合だった)かどうかを返す
func (uf *UnionFind) Unite(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)

	if rx == ry {
		return false
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
	uf.par[ry] = rx
	return true
}

// Same はxとyが同じノードにいるかを判断する
func (uf *UnionFind) Same(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)
	return rx == ry
}

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

// ch は condがtrueのときok, falseのときngを返します。
func ch(cond bool, ok, ng int) int {
	if cond {
		return ok
	}
	return ng
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
