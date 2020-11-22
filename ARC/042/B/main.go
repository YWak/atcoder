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

func main() {
	s := point{float64(nextInt()), float64(nextInt())}

	N := nextInt()
	p := make([]point, N)
	ans := float64(1000 * 1000)

	dist := func(x1, y1 float64) float64 {
		x, y := x1-s.x, y1-s.y
		return math.Sqrt(x*x + y*y)
	}
	distl := func(x1, y1, x2, y2 float64) float64 {
		if x1 == x2 {
			return math.Abs(x1 - s.x)
		}
		if y1 == y2 {
			return math.Abs(y1 - s.y)
		}

		// 直線
		x := x1 - x2
		y := y1 - y2

		a := y / x
		b := -1.0
		c := y1 - y/x*x1
		return math.Abs(a*s.x+b*s.y+c) / math.Sqrt(a*a+b*b)
	}

	for i := 0; i < N; i++ {
		x := float64(nextInt())
		y := float64(nextInt())
		p[i] = point{x, y}
		// 頂点までの距離
		ans = math.Min(dist(x, y), ans)
	}

	// 直線までの距離
	for i := 0; i < N; i++ {
		j := (i + 1) % N
		p1, p2 := p[i], p[j]
		d := distl(p1.x, p1.y, p2.x, p2.y)
		ans = math.Min(d, ans)
	}

	fmt.Printf("%.10f\n", ans)
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

// ==================================================
// 構造体
// ==================================================

// point は 座標を表す構造体です。
type point struct {
	x float64
	y float64
}
