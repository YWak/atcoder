//lint:file-ignore U1000 using template
package main

import (
	"bufio"
	"fmt"
	"io"
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

// N10_6は10^6
const N10_6 = int(1e6)

var in *In
var out *Out

func calc() {
	mod := 998244353
	n, r := 5000, 5000
	com := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		com[i] = make([]int, r+1)
	}
	com[0][0] = 1
	for i := 1; i <= n; i++ {
		com[i][0] = 1
		for j := 1; j <= r; j++ {
			com[i][j] = (com[i-1][j-1] + com[i-1][j]) % mod
		}
	}

	s := in.NextBytes()
	c := [26]int{}
	for _, v := range s {
		c[v-'a']++
	}
	// dp[i]は長さがjのときの場合の数
	dp := make([]int, len(s)+1)

	dp[0] = 1
	for _, l := range c {
		t := make([]int, len(s)+1)
		for k := 0; k <= l; k++ {
			for j := 0; j <= len(s); j++ {
				if j+k > len(s) {
					break
				}
				// j文字あるところにk文字挿入する
				pat := com[j+k][k]
				t[j+k] = (t[j+k] + pat*dp[j]%mod) % mod
			}
		}
		dp = t
	}
	ans := 0
	for i := 1; i <= len(s); i++ {
		ans = (ans + dp[i]) % mod
	}
	out.Println(ans)
}

type Combination interface {
	nCr(n, k int) int
}

type Combination1 struct {
	dp [][]int
}

func (c Combination1) nCr(n, r int) int {
	if n <= 0 || r <= 0 || n < r {
		return 1
	}
	return c.dp[n][r]
}

type Combination2 struct {
	mod   *Mod
	fact  []int
	ifact []int
}

func (c Combination2) nCr(n, r int) int {
	if n <= 0 || r <= 0 || n < r {
		return 1
	}
	return c.mod.mul(c.fact[n], c.mod.mul(c.ifact[r], c.ifact[n-r]))
}

type Combination3 struct {
	mod   *Mod
	ifact []int
}

func (c Combination3) nCr(n, r int) int {
	if n <= 0 || r <= 0 || n < r {
		return 1
	}
	ans := 1
	for i := n; i >= n-r+1; i-- {
		ans = c.mod.mul(ans, i)
	}
	return c.mod.mul(ans, c.ifact[r])
}

func NewComination(n, r int, mod *Mod) Combination {
	if n <= 5000 && r <= 5000 {
		// 完全に初期化できる
		dp := make([][]int, n+1)
		for i := 0; i <= n; i++ {
			dp[i] = make([]int, r+1)
		}
		dp[0][0] = 1
		for i := 1; i <= n; i++ {
			dp[i][0] = 1
			for j := 1; j <= r; j++ {
				dp[i][j] = mod.add(dp[i-1][j-1], dp[i-1][j])
			}
		}
		return &Combination1{dp}
	}

	initFact := func(m int) ([]int, []int) {
		N := n + 1
		// 全部の初期化が間に合う
		fact := make([]int, N)
		ifact := make([]int, N)
		fact[1] = 1
		for i := 2; i < N; i++ {
			fact[i] = mod.mul(fact[i-1], i)
		}
		ifact[n] = mod.inv(fact[n])
		for i := n; i > 0; i-- {
			ifact[i-1] = mod.mul(ifact[i], i)
		}
		return fact, ifact
	}

	if n <= pow(10, 7) {
		fact, ifact := initFact(n)
		return Combination2{mod, fact, ifact}
	}
	if r <= pow(10, 7) {
		_, ifact := initFact(r)
		return Combination3{mod, ifact}
	}

	panic("can not define")
}

type Mod struct {
	modulo int
}

func NewMod1000000007() *Mod {
	return &Mod{pow(10, 9) + 7}
}

func NewMod998244353() *Mod {
	return &Mod{998244353}
}

func (mod *Mod) norm(a int) int {
	if a < 0 || a >= mod.modulo {
		a %= mod.modulo
	}
	if a < 0 {
		a += mod.modulo
	}
	return a
}

func (mod *Mod) add(a, b int) int {
	ab := a + b
	if ab >= mod.modulo {
		ab %= mod.modulo
	}
	return ab
}

func (mod *Mod) sub(a, b int) int {
	ab := a - b + mod.modulo
	if ab >= mod.modulo {
		ab -= mod.modulo
	}
	return ab
}
func (mod *Mod) mul(a, b int) int {
	return (a * b) % mod.modulo
}

func (mod *Mod) pow(a, b int) int {
	ans := 1

	for b > 0 {
		if b&1 == 1 {
			ans = mod.mul(ans, a)
		}
		a = mod.mul(a, a)
		b = b >> 1
	}

	return ans
}

func (mod *Mod) inv(a int) int {
	// 拡張ユークリッドの互除法
	b, u, v := mod.modulo, 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	return mod.norm(u)
}

func (mod *Mod) div(a, b int) int {
	return mod.mul(a, mod.inv(b))
}

func (mod *Mod) chadd(a *int, b int) {
	*a = mod.add(*a, b)
}

func (mod *Mod) chsub(a *int, b int) {
	*a = mod.sub(*a, b)
}

func (mod *Mod) chmul(a *int, b int) {
	*a = mod.mul(*a, b)
}

func (mod *Mod) chdiv(a *int, b int) {
	*a = mod.div(*a, b)
}

func main() {
	// interactiveならfalseにすること。
	in, out = InitIo(true)
	defer out.Flush()

	calc()
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

// ==================================================
// 入出力操作
// ==================================================
type In struct {
	// NextString は 次の入力を文字列として読み込んで返します。
	NextString func() string
}

type Out struct {
	writer io.Writer
	Flush  func()
}

// InitIo は inとoutを初期化します。
func InitIo(buffer bool) (*In, *Out) {
	bufsize := 4 * 1024 * 1024 // 4MB

	// 入力はずっとバッファーでいいらしい。ほんとう？
	// TODO バッファなしfmt.Fscanf(os.Stdin)だとTLEだった。要調査
	_in := bufio.NewScanner(os.Stdin)
	_in.Split(bufio.ScanWords)
	_in.Buffer(make([]byte, bufsize), bufsize)
	in := func() string {
		_in.Scan()
		return _in.Text()
	}

	// 出力はバッファon/offが必要
	var out io.Writer
	var flush func()

	if buffer {
		_out := bufio.NewWriterSize(os.Stdout, bufsize)
		out = _out
		flush = func() {
			_out.Flush()
		}
	} else {
		out = os.Stdout
		flush = func() {}
	}

	return &In{in}, &Out{out, flush}
}

// NextBytes は 次の入力をbyteの配列として読み込んで返します。
// 遅いから極力使わない。
func (in *In) NextBytes() []byte {
	return []byte(in.NextString())
}

// NextInt は 次の入力を数値として読み込んで返します。
func (in *In) NextInt() int {
	i, _ := strconv.Atoi(in.NextString())
	return i
}

// NextInt2 は 次の2つの入力を数値として読み込んで返します。
func (in *In) NextInt2() (int, int) {
	return in.NextInt(), in.NextInt()
}

// NextInt2d は 次の2つの入力を数値n1,n2として読み込んで、n1+d1, n2+d2を返します。
func (in *In) NextInt2d(d1, d2 int) (int, int) {
	return in.NextInt() + d1, in.NextInt() + d2
}

// NextInt3 は 次の3つの入力を数値として読み込んで返します。
func (in *In) NextInt3() (int, int, int) {
	return in.NextInt(), in.NextInt(), in.NextInt()
}

// NextInt2d は 次の3つの入力を数値n1,n2,n3として読み込んで、n1+d1, n2+d2, n3+d3を返します。
func (in *In) NextInt3d(d1, d2, d3 int) (int, int, int) {
	return in.NextInt() + d1, in.NextInt() + d2, in.NextInt() + d3
}

// NextInt4 は 次の4つの入力を数値として読み込んで返します。
func (in *In) NextInt4() (int, int, int, int) {
	return in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
}

// NextInts は 次のn個の入力を数値として読み込んで、配列として返します。
func (in *In) NextInts(n int) sort.IntSlice {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = in.NextInt()
	}
	return sort.IntSlice(a)
}

// NextLongIntAsArray は 次の入力を数値として読み込み、各桁を要素とした配列を返します。
func (in *In) NextLongIntAsArray() []int {
	s := in.NextString()
	l := len(s)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = int(s[i] - '0')
	}

	return arr
}

// NextFloat は 次の入力を実数値として読み込み、値を返します。
func (in *In) NextFloat() float64 {
	f, _ := strconv.ParseFloat(in.NextString(), 64)
	return f
}

// NextFloatAsInt は次の入力を実数rとして読み込み、r * 10^base の値を返します。
func (in *In) NextFloatAsInt(base int) int {
	if base%10 == 0 {
		panic("baseは小数点の最大桁数を指定する")
	}

	s := in.NextString()
	index := strings.IndexByte(s, '.')

	// 小数点がなければそのまま返す
	if index == -1 {
		n, _ := strconv.Atoi(s)
		return n * pow(10, base)
	}

	// 末尾の0を消しておく
	for s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}

	// 整数部分 * 10^base + 小数部分 * 10^(足りない分)
	s1, s2 := s[:index], s[index+1:]
	n, _ := strconv.Atoi(s1)
	m, _ := strconv.Atoi(s2)

	return n*pow(10, base) + m*pow(10, base-len(s2))
}

// Println は引数をスペース区切りで出力し、最後に改行を出力します。
func (out *Out) Println(a ...interface{}) {
	fmt.Fprintln(out.writer, a...)
}

// Printf はformatにしたがってaを整形して出力します。
func (out *Out) Printf(format string, a ...interface{}) {
	fmt.Fprintf(out.writer, format, a...)
}

// PrintStringsln は文字列配列の各要素をスペース区切りで出力し、最後に改行を出力します。
func (out *Out) PrintStringsln(a []string) {
	out.Println(strings.Join(a, " "))
}

// PrintIntsLn は整数配列の各要素をスペース区切りで出力し、最後に改行を出力します。
func (out *Out) PrintIntsLn(a []int) {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = fmt.Sprint(v)
	}
	out.Println(strings.Join(b, " "))
}

func (out *Out) PrintLenAndIntsLn(a []int) {
	b := make([]string, len(a)+1)
	b[0] = fmt.Sprint(len(a))
	for i, v := range a {
		b[i+1] = fmt.Sprint(v)
	}
	out.Println(strings.Join(b, " "))
}

// Putcは一文字出力します。
func (out *Out) Putc(c byte) {
	out.Printf("%c", c)
}

// YesNo は condが真ならYes, 偽ならNoを出力します。
func (out *Out) YesNo(cond bool) {
	if cond {
		out.Println("Yes")
	} else {
		out.Println("No")
	}
}

func (out *Out) YESNO(cond bool) {
	if cond {
		out.Println("YES")
	} else {
		out.Println("NO")
	}
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
	ans := 1
	for b > 0 {
		if b%2 == 1 {
			ans *= a
		}
		a, b = a*a, b/2
	}
	return ans
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
	if n == 0 {
		return 1
	}

	x = x % m
	if x == 0 {
		return 0
	}

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
// ==================================================
// NewIntInt は数値の二次元配列を作成します。
func NewIntInt(rows, cols, val int) [][]int {
	a := make([][]int, rows)
	for i := 0; i < rows; i++ {
		a[i] = make([]int, cols)

		for j := 0; j < cols; j++ {
			a[i][j] = val
		}
	}

	return a
}

// reverseIntはintの配列を逆転させます。
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

// compressはnumbersで渡した値を座標圧縮します。
func compress(numbers map[int]int) (map[int]int, []int) {
	keys := sort.IntSlice{}
	for i := range numbers {
		keys = append(keys, i)
	}
	sort.Sort(keys)
	for i, v := range keys {
		numbers[v] = i
	}

	return numbers, keys
}
