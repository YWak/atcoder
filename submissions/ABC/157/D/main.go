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

func main() {
	N := nextInt()
	M := nextInt()
	K := nextInt()

	group := ufNew(N)
	friends := make([]int, N)
	blocks := make([]int, N)

	for i := 0; i < M; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		group.Unite(a, b)
		friends[a]++
		friends[b]++
	}
	for i := 0; i < K; i++ {
		c := nextInt() - 1
		d := nextInt() - 1
		if group.Same(c, d) {
			blocks[c]++
			blocks[d]++
		}
	}
	ans := make([]string, N)
	for i := 0; i < N; i++ {
		ans[i] = strconv.Itoa(group.Size(i) - friends[i] - blocks[i] - 1)
	}
	fmt.Println(strings.Join(ans, " "))
}

// UnionFind : UnionFind構造を保持する構造体
type UnionFind struct {
	par  []int // i番目のノードに対応する親
	rank []int // i番目のノードのランク
	size []int // i番目のノードに紐づくノードの数
}

// [0, n)のノードを持つUnion-Findを作る
func ufNew(n int) UnionFind {
	uf := UnionFind{par: make([]int, n), rank: make([]int, n), size: make([]int, n)}

	for i := 0; i < n; i++ {
		uf.par[i] = i
		uf.size[i] = 1
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
	s := uf.size[rx] + uf.size[ry]
	uf.size[rx] = s
	uf.size[ry] = s
	uf.par[ry] = rx
	return true
}

// Same はxとyが同じノードにいるかを判断する
func (uf *UnionFind) Same(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)
	return rx == ry
}

// Size はxのグループに所属するノードの数
func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Root(x)]
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
