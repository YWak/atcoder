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

var dp = [101][101][101]float64{}

func main() {
	A := nextInt()
	B := nextInt()
	C := nextInt()

	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			for k := 0; k < 101; k++ {
				if i == 100 || j == 100 || k == 100 {
					dp[i][j][k] = 0
				} else {
					dp[i][j][k] = -1
				}
			}
		}
	}

	fmt.Printf("%.9f\n", solve(A, B, C))
}

func solve(a, b, c int) float64 {
	ans := dp[a][b][c]
	if ans == -1 {
		s := float64(a + b + c)

		ans = float64(a)/s*(1+solve(a+1, b, c)) + float64(b)/s*(1+solve(a, b+1, c)) + float64(c)/s*(1+solve(a, b, c+1))
		dp[a][b][c] = ans
	}

	return ans
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
