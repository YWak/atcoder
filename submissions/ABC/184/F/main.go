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
	T := nextInt()
	A := nextInts(N)

	m := N / 2
	n := N - m
	heads := make([]int, pow(2, m))
	tails := make([]int, pow(2, n))

	for i := 0; i < (1 << m); i++ {
		for j := 0; j < m; j++ {
			if nthbit(i, j) == 1 {
				heads[i] += A[j]
			}
		}
	}
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			if nthbit(i, j) == 1 {
				tails[i] += A[m+j]
			}
		}
	}
	sort.Slice(tails, func(i, j int) bool {
		return tails[i] < tails[j]
	})

	ans := 0
	for i := 0; i < len(heads); i++ {
		j := binarysearch(-1, len(tails), func(n int) bool {
			return tails[n] <= T-heads[i]
		})
		if j < 0 || j >= len(tails) {
			continue
		}
		ans = max(ans, heads[i]+tails[j])
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
