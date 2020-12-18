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
	Q := nextInt()
	A := make(SegmentTree, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	A.Init()

	for i := 0; i < Q; i++ {
		t := nextInt()
		x := nextInt()
		y := nextInt()

		if t == 1 {
			A.Update(x-1, y)
		} else {
			fmt.Println(A.Query(x-1, y))
		}
	}
}

type SegmentTree []int

func (st *SegmentTree) Init() {
	n := len(*st)
	x := 1
	for x < n {
		x *= 2
	}

	arr := make(SegmentTree, x*2)
	for i := 0; i < x*2-1; i++ {
		arr[i] = st.E()
	}
	for i := 0; i < n; i++ {
		arr.Update(i, (*st)[i])
	}

	*st = arr
}

func (st SegmentTree) E() int {
	return 0
}
func (st SegmentTree) Operate(a, b int) int {
	return a ^ b
}

func (st SegmentTree) Update(i int, value int) {
	i += len(st)/2 - 1
	st[i] ^= value
	for i > 0 {
		i = (i - 1) / 2
		st[i] = st.Operate(st[i*2+1], st[i*2+2])
	}
}

func (st SegmentTree) Query(from, to int) int {
	return st.query1(from, to, 0, 0, len(st)/2)
}

func (st SegmentTree) query1(from, to, k, l, r int) int {
	if r <= from || to <= l {
		return st.E()
	}
	if from <= l && r <= to {
		return st[k]
	}
	m := (l + r) / 2
	vl := st.query1(from, to, k*2+1, l, m)
	vr := st.query1(from, to, k*2+2, m, r)

	return st.Operate(vl, vr)
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

// nextInt は 入力から整数をひとつ読み込みます
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
