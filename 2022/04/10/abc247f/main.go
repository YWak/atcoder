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

var in *In
var out *Out

func calc() {
	n := in.NextInt()
	p := in.NextInts(n)
	q := in.NextInts(n)

	uf := ufNew(n)
	for i := 0; i < n; i++ {
		uf.Unite(p[i]-1, q[i]-1)
	}

	mod := NewMod998244353()
	c := NewComination(200000, 200000, mod)
	ans := 1
	used := map[int]bool{}

	for i := 0; i < n; i++ {
		root := uf.Root(i)
		if used[root] {
			continue
		}
		used[root] = true
		s := uf.Size(root)

		if s == 1 {
			ans = mod.mul(ans, 1)
		} else if s == 3 {
			ans = mod.mul(ans, 4) // どれかひとつ選ばない/全部
		} else if s%2 == 1 {
			k := 0
			ss := s/2 - 1
			for j := 0; j <= ss; j++ {
				k = mod.add(k, c.nCk(ss, j))
			}
			k = mod.mul(k, 2)
			k = mod.sub(k, 1)
			ans = mod.mul(ans, k)
		} else {
			// 半数選んだところから追加でt個選べる
			k := 0
			for j := 0; j <= s/2; j++ {
				k = mod.add(k, c.nCk(s/2, j))
			}

			k = mod.mul(k, 2) //
			k = mod.sub(k, 1) // 全部選ぶ分がかぶる
			ans = mod.mul(ans, k)
		}
	}

	out.Println(ans)
}

type Combination struct {
	mod   *Mod
	style int
	fact  []int
	ifact []int
	dp    [][]int
}

func NewComination(n, k int, mod *Mod) Combination {
	c := Combination{mod: mod}

	if n <= 5000 && k <= 5000 {
		c.style = 1
		// 完全に初期化できる
		c.dp = make([][]int, n+1)
		for i := 0; i <= n; i++ {
			c.dp[i] = make([]int, k+1)
		}
		c.dp[0][0] = 1
		for i := 1; i <= n; i++ {
			c.dp[i][0] = 1
			for j := 1; j <= k; j++ {
				c.dp[i][j] = c.mod.add(c.dp[i-1][j-1], c.dp[i-1][j])
			}
		}
	} else if n <= pow(10, 7) {
		c.style = 2
		c.initFact(n)
	} else if k <= pow(10, 7) {
		c.style = 3
		c.initFact(k)
	}
	return c
}

func (c *Combination) initFact(n int) {
	N := n + 1
	// 全部の初期化が間に合う
	c.fact = make([]int, N)
	c.ifact = make([]int, N)
	c.fact[1] = 1
	for i := 2; i < N; i++ {
		c.fact[i] = c.mod.mul(c.fact[i-1], i)
	}
	c.ifact[n] = c.mod.inv(c.fact[n])
	for i := n; i > 0; i-- {
		c.ifact[i-1] = c.mod.mul(c.ifact[i], i)
	}
}

func (c *Combination) nCk(n, k int) int {
	switch c.style {
	case 1:
		return c.dp[n][k]
	case 2:
		return c.mod.mul(c.fact[n], c.mod.mul(c.ifact[k], c.ifact[n-k]))
	case 3:
		ans := 1
		for i := n; i >= n-k+1; i-- {
			ans = c.mod.mul(ans, i)
		}
		return c.mod.mul(ans, c.ifact[k])
	default:
		panic("not initialized")
	}
}

type Mod struct {
	modulo int

	// normはaをmod mの値に変換します
	norm func(a int) int

	// addはa + b (mod m)を返します。
	add func(a, b int) int

	// subはa - b (mod m)を返します。
	sub func(a, b int) int

	// mulはa * b (mod m)を返します。
	mul func(a, b int) int

	// powはa ^ b (mod m)を返します。
	pow func(a, b int) int

	// invはmod mにおけるaの逆元を返します。
	inv func(a int) int

	// divはa / b (mod m)を返します。
	div func(a, b int) int
}

func NewMod1000000007() *Mod {
	mod := NewMod(1000000007)
	return mod
}

func NewMod998244353() *Mod {
	mod := NewMod(998244353)
	return mod
}

func NewMod(m int) *Mod {
	norm := func(a int) int {
		if a < 0 || a >= m {
			a %= m
		}
		if a < 0 {
			a += m
		}
		return a
	}
	add := func(a, b int) int {
		ab := a + b
		if ab >= m {
			ab %= m
		}
		return ab
	}
	sub := func(a, b int) int {
		ab := a - b + m
		if ab < 0 {
			ab += m
		}
		return ab
	}
	mul := func(a, b int) int {
		return (a * b) % m
	}
	pow := func(a, b int) int {
		ans := 1

		for b > 0 {
			if b&1 == 1 {
				ans = mul(ans, a)
			}
			a = mul(a, a)
			b = b >> 1
		}

		return ans
	}
	inv := func(a int) int {
		// 拡張ユークリッドの互除法
		b, u, v := m, 1, 0
		for b > 0 {
			t := a / b
			a -= t * b
			a, b = b, a
			u -= t * v
			u, v = v, u
		}
		return norm(u)
	}
	div := func(a, b int) int {
		return mul(a, inv(b))
	}

	return &Mod{
		modulo: m,
		norm:   norm,
		add:    add,
		sub:    sub,
		mul:    mul,
		pow:    pow,
		inv:    inv,
		div:    div,
	}
}

// UnionFind : UnionFind構造を保持する構造体
type UnionFind struct {
	par []int // i番目のノードに対応する親
}

// [0, n)のノードを持つUnion-Findを作る
func ufNew(n int) UnionFind {
	uf := UnionFind{par: make([]int, n)}

	for i := 0; i < n; i++ {
		uf.par[i] = -1
	}

	return uf
}

// Root はxのルートを得る
func (uf *UnionFind) Root(x int) int {
	if uf.par[x] < 0 {
		return x
	}
	uf.par[x] = uf.Root(uf.par[x])
	return uf.par[x]
}

// Unite はxとyを併合する。集合の構造が変更された(== 呼び出し前は異なる集合だった)かどうかを返す
func (uf *UnionFind) Unite(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)

	if rx == ry {
		return false
	}
	if uf.par[rx] > uf.par[ry] {
		rx, ry = ry, rx
	}
	uf.par[rx] += uf.par[ry]
	uf.par[ry] = rx
	return true
}

// Same はxとyが同じノードにいるかを判断する
func (uf *UnionFind) Same(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)
	return rx == ry
}

// Size は xの集合のサイズを返します。
func (uf *UnionFind) Size(x int) int {
	return -uf.par[uf.Root(x)]
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

// YesNo は condが真ならYes, 偽ならNoを出力します。
func (out *Out) YesNo(cond bool) {
	if cond {
		out.Println("Yes")
	} else {
		out.Println("No")
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
