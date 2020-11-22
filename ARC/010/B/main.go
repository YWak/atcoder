package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	N := nextInt()
	dates := [800]bool{}
	for i := 0; i < 367; i++ {
		dates[i] = (i%7) == 0 || (i%7) == 6
	}
	holidays := make([]date, N)
	for i := 0; i < N; i++ {
		d := toDate(nextString())
		holidays[i] = d
	}
	sort.Slice(holidays, func(i, j int) bool {
		h := holidays
		return h[i].m < h[j].m || h[i].m == h[i].m && h[i].d < h[j].d
	})
	for i := 0; i < N; i++ {
		n := toNum(holidays[i])
		for dates[n] {
			n++
		}
		dates[n] = true
	}
	ans := 0
	c := 0

	dates[366] = false
	for i := 0; i < 367; i++ {
		if dates[i] {
			c++
		} else {
			ans = max(ans, c)
			c = 0
		}
	}

	fmt.Println(ans)
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func toNum(d date) int {
	days := []int{-1, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	n := d.d - 1
	for i := 1; i < d.m; i++ {
		n += days[i]
	}

	return n
}

func toDate(s string) date {
	m := 0
	d := 0

	slash := false

	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			slash = true
		} else if slash {
			d = d*10 + toi(s[i])
		} else {
			m = m*10 + toi(s[i])
		}
	}
	return date{m, d}
}

type date struct {
	m int
	d int
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
