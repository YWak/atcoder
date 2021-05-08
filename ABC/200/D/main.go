//lint:file-ignore U1000 using template
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
const INF18 = int(1e18)

// INF9 は最大値を表す数
const INF9 = int(1e9)

func calc() {
	n := nextInt()
	a := nextInts(n)

	// dp[i][j] はi番目まで使ってあまりがjになる場合の数
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 200)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 200; j++ {
			dp[i+1][(j+a[i])%200] += dp[i][j]
			dp[i+1][j] += dp[i][j]
		}
	}
	s := -1
	for j := 0; j < 200; j++ {
		if dp[n][j] > 1 {
			s = j
			break
		}
	}
	if s == -1 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	debug(s)

	b := make([]bool, n)
	c := make([]bool, n)

	// 分岐するまでは同じになる
	var jj int
	var sb int
	var sc int
	cb := 0
	cc := 0
	mod := func(k int) int {
		return (200 + k) % 200
	}
	for j := n - 1; j >= 0; j-- {
		if dp[j+1][s] == dp[j][s] {
			// 両方使わない
			b[j] = false
			c[j] = false
		} else if dp[j][mod(s-a[j])] == dp[j+1][s] {
			// 両方使う
			b[j] = true
			c[j] = true
			cb++
			cc++
			s = mod(s - a[j])
		} else {
			b[j] = true
			c[j] = false
			cb++
			sb = mod(s - a[j])
			sc = s
			jj = j - 1
			break
		}
	}
	// debug("mid", jj, sb, sc, b, c)
	// for i := 0; i < len(dp); i++ {
	// 	m := map[int]int{}
	// 	for j := 0; j < len(dp[i]); j++ {
	// 		if dp[i][j] != 0 {
	// 			m[j] += dp[i][j]
	// 		}
	// 	}
	// 	debug(m)
	// }
	// bの復元
	for j := jj; j >= 0; j-- {
		if dp[j+1][mod(sb-a[j])] != 0 {
			b[j] = true
			sb = mod(sb - a[j])
			cb++
		}
	}
	// cの復元
	for j := jj; j >= 0; j-- {
		if dp[j+1][mod(sc-a[j])] != 0 {
			c[j] = true
			sc = mod(sc - a[j])
			cc++
		}
	}

	ansb := 0
	ansc := 0
	fmt.Print(cb)
	for i := 0; i < n; i++ {
		if b[i] {
			fmt.Printf(" %d", i+1)
			ansb += a[i]
		}
	}
	fmt.Println()
	fmt.Print(cc)
	for i := 0; i < n; i++ {
		if c[i] {
			fmt.Printf(" %d", i+1)
			ansc += a[i]
		}
	}
	fmt.Println()
	// debug("sum", ansb%200, ansc%200)
}

func main() {
	calc()
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

// divceil は a/b の結果を正の無限大に近づけるように丸めて返します。
func divceil(a, b int) int {
	if a%b == 0 || a/b < 0 {
		return a / b
	}
	return (a + b - 1) / b
}

// divfloor は a/b の結果を負の無限大に近づけるように丸めて返します。
func divfloor(a, b int) int {
	if a%b == 0 || a/b > 0 {
		return a / b
	}
	if b < 0 {
		a, b = -a, -b
	}
	return (a - b + 1) / b
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
