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
	s := point{nextInt(), nextInt()}
	g := point{nextInt(), nextInt()}

	c := 3

	if s == g {
		c = 0
	} else if canMove(s, g) {
		c = 1
	} else if canMove2(s, g) {
		c = 2
	} else {
		for i := -3; i <= 3; i++ {
			for j := -3; j <= 3; j++ {
				if abs(i)+abs(j) > 3 {
					continue
				}

				g1 := point{g.x + i, g.y + j}
				if abs(s.x-g1.x) == abs(s.y-g1.y) {
					c = 2
				}
			}
		}
	}

	fmt.Println(c)
}

func canMove(p1, p2 point) bool {
	if p1.x+p1.y == p2.x+p2.y {
		return true
	}
	if p1.x-p1.y == p2.x-p2.y {
		return true
	}
	if abs(p1.x-p2.x)+abs(p1.y-p2.y) <= 3 {
		return true
	}
	return false
}

func canMove2(p1, p2 point) bool {
	if (p2.x+p2.y-p1.x-p1.y)%2 != 0 {
		return false
	}
	if (p2.x-p1.x-p2.y+p1.y)%2 != 0 {
		return false
	}
	return true
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
