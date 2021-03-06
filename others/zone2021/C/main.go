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

type man struct{ i, a, b, c, d, e int }

func ywak() {
	n := nextInt()
	mm := make([]man, 0, n)
	ma := make([]man, 0, n)
	mb := make([]man, 0, n)
	mc := make([]man, 0, n)
	md := make([]man, 0, n)
	me := make([]man, 0, n)

	for i := 0; i < n; i++ {
		a, b, c, d, e := nextInt5()
		m := man{i, a, b, c, d, e}
		mm = append(mm, m)
		ma = append(ma, m)
		mb = append(mb, m)
		mc = append(mc, m)
		md = append(md, m)
		me = append(me, m)
	}
	sort.Slice(ma, func(i, j int) bool { return ma[i].a > ma[j].a })
	sort.Slice(mb, func(i, j int) bool { return mb[i].b > mb[j].b })
	sort.Slice(mc, func(i, j int) bool { return mc[i].c > mc[j].c })
	sort.Slice(md, func(i, j int) bool { return md[i].d > md[j].d })
	sort.Slice(me, func(i, j int) bool { return me[i].e > me[j].e })

	ans := 0
	min5 := func(a, b, c, d, e int) int { return min(a, min(b, min(c, min(d, e)))) }
	max3 := func(a, b, c int) int { return max(a, max(b, c)) }
	score := func(m1, m2, m3 man) int {
		a := max3(m1.a, m2.a, m3.a)
		b := max3(m1.b, m2.b, m3.b)
		c := max3(m1.c, m2.c, m3.c)
		d := max3(m1.d, m2.d, m3.d)
		e := max3(m1.e, m2.e, m3.e)

		return min5(a, b, c, d, e)
	}
	find := func(a, b int, mz []man) int {
		for i := 0; i < 3; i++ {
			if mz[i].i != a && mz[i].i != b {
				return i
			}
		}
		return -1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, score(mm[i], mm[j], ma[find(i, j, ma)]))
			ans = max(ans, score(mm[i], mm[j], mb[find(i, j, mb)]))
			ans = max(ans, score(mm[i], mm[j], mc[find(i, j, mc)]))
			ans = max(ans, score(mm[i], mm[j], md[find(i, j, md)]))
			ans = max(ans, score(mm[i], mm[j], me[find(i, j, me)]))
		}
	}

	fmt.Println(ans)
}

func chokudai() {
	K := 5
	M := 1 << K
	n := nextInt()
	best := make([]int, M)

	for i := 0; i < n; i++ {
		m := nextInts(K)

		for j := 0; j < M; j++ {
			t := INF9
			for k := 0; k < K; k++ {
				if (j>>k)&1 == 1 {
					t = min(t, m[k])
				}
			}
			best[j] = max(best[j], t)
		}
	}

	ans := 0
	for i := 0; i < M; i++ {
		for j := i; j < M; j++ {
			for k := j; k < M; k++ {
				if i^j^k != M-1 {
					continue
				}
				a := min(best[i], min(best[j], best[k]))
				ans = max(ans, a)
			}
		}
	}
	fmt.Println(ans)
}

func main() {
	chokudai()
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

func nextInt5() (int, int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
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
