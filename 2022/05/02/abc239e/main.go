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
	n, Q := in.NextInt2()
	x := in.NextInts(n)

	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := in.NextInt2d(-1, -1)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	l := make([]int, n)
	r := make([]int, n)
	arr := []int{}
	t := 0
	var dfs func(curr, prev int)
	dfs = func(curr, prev int) {
		arr = append(arr, x[curr])
		l[curr] = t
		t++
		for _, next := range g[curr] {
			if next == prev {
				continue
			}
			dfs(next, curr)
		}
		r[curr] = t
	}
	dfs(0, -1)

	w := NewWaveletMatrix(arr)
	for q := 0; q < Q; q++ {
		v, k := in.NextInt2d(-1, -1)
		out.Println(w.KthLargest(l[v], r[v], k))
	}
}

type Dictionary []int

func (d Dictionary) rank0(r int) int {
	return r - d[r]
}

func (d Dictionary) rank1(r int) int {
	return d[r]
}

type WaveletMatrix struct {
	values  []int
	index   [][]int
	zeros   []int
	sum     []Dictionary
	start   map[int]int
	bitsize int
}

// NewWaveletMatrixはarrをもとにしてWaveletMatrixの実装を返します。
func NewWaveletMatrix(arr []int) *WaveletMatrix {
	n := len(arr)
	values1 := make([]int, n)
	values2 := make([]int, n)

	// 1である最大のbitを取得する
	bitsize := 0
	for i, v := range arr {
		values1[i] = v
		for b := 0; b < 64; b++ {
			if (v >> b) > 0 {
				bitsize = max(bitsize, b+1)
			}
		}
	}

	index := make([][]int, bitsize)
	sum := make([]Dictionary, bitsize)
	zeros := make([]int, bitsize)

	for b := bitsize - 1; b >= 0; b-- {
		// インデックスと累積和の更新
		index[b] = make([]int, n)
		sum[b] = make(Dictionary, n+1)
		for i, v := range values1 {
			t := (v >> b) & 1
			if t == 0 {
				zeros[b]++
			}
			index[b][i] = t
			sum[b][i+1] = sum[b][i] + t
		}

		// 安定ソート
		k := 0
		for i := 0; i < n; i++ {
			if index[b][i] == 0 {
				values2[k] = values1[i]
				k++
			}
		}
		for i := 0; i < n; i++ {
			if index[b][i] == 1 {
				values2[k] = values1[i]
				k++
			}
		}
		values1, values2 = values2, values1
	}
	start := map[int]int{}
	for i := n - 1; i >= 0; i-- {
		start[values1[i]] = i
	}
	return &WaveletMatrix{values1, index, zeros, sum, start, bitsize}
}

// Getはk(0-indexed)番目の要素を返します。
func (w *WaveletMatrix) Get(k int) int {
	v := 0
	i := k
	for b := w.bitsize - 1; b >= 0; b-- {
		t := w.index[b][i]
		v += t << b
		if t == 0 {
			i = w.sum[b].rank0(i)
		} else {
			i = w.zeros[b] + w.sum[b].rank1(i)
		}
	}

	return v
}

// Rankは区間[0, r)に含まれるxの個数を返します。
// tested:
//   https://atcoder.jp/contests/abc248/tasks/abc248_d
func (w *WaveletMatrix) Rank(x, r int) int {
	if x > (1 << (w.bitsize + 1)) {
		return 0
	}

	i := r
	for b := w.bitsize - 1; b >= 0; b-- {
		if (x>>b)&1 == 0 {
			i = w.sum[b].rank0(i)
		} else {
			i = w.zeros[b] + w.sum[b].rank1(i)
		}
	}

	return i - w.start[x]
}

// KthSmallestは区間[l, r)に含まれる要素のうちk番目(0-indexed)に小さいものを返します。
// tested:
//   https://judge.yosupo.jp/problem/range_kth_smallest
func (w *WaveletMatrix) KthSmallest(l, r, k int) int {
	if l < 0 || l >= len(w.index[0]) || r < 0 || r > len(w.index[0]) {
		panic(fmt.Sprintf("invalid range [%d, %d)", l, r))
	}

	v := 0
	for b := w.bitsize - 1; b >= 0; b-- {
		c0 := w.sum[b].rank0(r) - w.sum[b].rank0(l)

		if c0 > k {
			// 0の行き先を探す
			l = w.sum[b].rank0(l)
			r = w.sum[b].rank0(r)
		} else {
			// 1の行き先を探す
			l = w.zeros[b] + w.sum[b].rank1(l)
			r = w.zeros[b] + w.sum[b].rank1(r)
			k = k - c0
			v += (1 << b)
		}
	}

	return v
}

// KthLargestは区間[l, r)に含まれる要素のうちk番目(0-indexed)に大きいものを返します。
// tested:
//   https://atcoder.jp/contests/abc239/tasks/abc239_e
func (w *WaveletMatrix) KthLargest(l, r, k int) int {
	return w.KthSmallest(l, r, r-l-1-k)
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
