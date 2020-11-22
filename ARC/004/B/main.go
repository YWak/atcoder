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
	N := nextInt()
	A := nextInts(N)

	sum := 0
	heads := make([]int, N)
	tails := make([]int, N+1)

	for i := 0; i < N; i++ {
		sum += A[i]
	}
	fmt.Println(sum)

	heads[0] = A[0]
	for i := 1; i < N; i++ {
		heads[i] = heads[i-1] + A[i]
	}

	for i := N - 1; i >= 0; i-- {
		tails[i] = tails[i+1] + A[i]
	}

	if N == 2 {
		fmt.Println(abs(A[0] - A[1]))
		return
	}
	ans := sum
	for i := 1; i < N-1; i++ {
		a := heads[i-1]
		b := A[i]
		c := tails[i+1]

		if b < a+c && a < b+c && c < a+b {
			// 三角になる
			ans = 0
		} else {
			n := 0
			n = max(n, b-a-c)
			n = max(n, a-b-c)
			n = max(n, c-a-b)
			ans = min(ans, n)
		}
	}
	fmt.Println(ans)
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
	x int
	y int
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
