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

// INF18 は最大値を表す数
const INF18 = int(1e18) * 2

// INF9 は最大値を表す数
const INF9 = int(1e9)

// crtは中国剰余定理の実装です。
// x = b[i] mod m[i]
// となる x mod lcm(m) を計算し、 x, lcm(m) を返します。
// 解なしの場合は 0, 0を返します。
func crt(b, m []int) (int, int) {
	b0, m0 := 0, 1
	invgcd := func(a, m int) (int, int) {
		x, u := 1, 0
		for m != 0 {
			t := a / m
			a, m = m, a-t*m
			x, u = u, x-t*u
		}
		return a, x
	}
	for i := 0; i < len(b); i++ {
		b1, m1 := b[i], m[i]
		if m0 < m1 {
			b0, b1 = b1, b0
			m0, m1 = m1, m0
		}
		if m0%m1 == 0 {
			if b0%m1 != b1 {
				return 0, 0
			}
			continue
		}

		g, im := invgcd(m0, m1)
		if (b1-b0)%g != 0 {
			return 0, 0
		}
		u := m1 / g
		x := (b1 - b0) / g % u * im % u
		b0 += m0 * x
		m0 *= u
	}
	return (b0%m0 + m0) % m0, m0
}

func solve(X, Y, P, Q int) {
	ans := INF18

	s := 2 * (X + Y)
	t := P + Q

	for y := 0; y < Y; y++ {
		for q := 0; q < Q; q++ {
			a, b := crt([]int{X + y, P + q}, []int{s, t})
			if b == 0 {
				continue
			}
			ans = min(ans, a)
		}
	}

	if ans == INF18 {
		fmt.Println("infinity")
	} else {
		fmt.Println(ans)
	}
}

func main() {
	t := nextInt()
	for i := 0; i < t; i++ {
		x, y, p, q := nextInt4()
		solve(x, y, p, q)
	}
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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
}

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
}

func nextInts(n int) sort.IntSlice {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return sort.IntSlice(a)
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

// nextFloatAsInt は 数を 10^base 倍した整数値を取得します。
func nextFloatAsInt(base int) int {
	s := nextString()
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

// ceil は a/bの切り上げを返します。
func ceil(a, b int) int {
	return (a + b - 1) / b
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
