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

var n int
var m int

type query struct {
	op, l, r, x, i, j int
}

func index(i, j int) int {
	return j*m + i
}

func calc() {
	n, m = in.NextInt2()
	Q := in.NextInt()

	st := NewSegmentTree()
	n2i := map[int]int{}

	queries := make([]*query, Q)
	last := make([]int, n)       // i行を最後に参照したop = 2のクエリ番号
	watch := make([][]*query, Q) // op = 2を参照する op = 3のクエリのリスト
	for k := 0; k < Q; k++ {
		op := in.NextInt()
		var q query
		switch op {
		case 1:
			{
				l, r, x := in.NextInt3()
				l--
				q = query{op: op, l: l, r: r, x: x}
				n2i[index(0, l)] = 1
				n2i[index(n-1, r-1)] = 1
				n2i[index(0, r)] = 1
			}
		case 2:
			{
				i, x := in.NextInt2()
				i--
				q = query{op: op, i: i, x: x}
				last[i] = k // i行目に最後にop = 2の操作をしたタイミングを記録しておく
			}
		case 3:
			{
				i, j := in.NextInt2()
				i--
				j--
				q = query{op: op, i: i, j: j}
				n2i[index(i, j)] = 1
				watch[last[i]] = append(watch[last[i]], &q)
			}
		}
		queries[k] = &q
	}
	compress1(n2i)
	st.init(len(n2i) + 1)

	for i, q := range queries {
		switch q.op {
		case 1:
			st.update(n2i[index(0, q.l)], n2i[index(0, q.r)], q.x)
		case 2:
			for _, v := range watch[i] {
				v.x = q.x - st.query(n2i[index(v.i, v.j)])
			}
		case 3:
			out.Println(q.x + st.query(n2i[index(q.i, q.j)]))
		}
	}
}

type SegmentTreeFunctions struct {
	// 単位元を返します
	e func() int
	// 計算結果を返します
	calc func(a, b int) int
}

type SegmentTree struct {
	// このsegment treeが管理するインデックスの範囲。[0, n)を管理する。
	n int

	// segment treeの各ノードの値を保持する配列
	nodes []int

	// このsegment treeの値を操作する関数群
	f SegmentTreeFunctions
}

// NewSegmentTreeは区間和を扱うSegmentTreeを返します。
// tested:
//   https://atcoder.jp/contests/abl/tasks/abl_d
func NewSegmentTree() *SegmentTree {
	return &SegmentTree{
		-1,
		[]int{},
		SegmentTreeFunctions{
			func() int { return 0 },
			func(a, b int) int { return a + b },
		},
	}
}

// NewRangeMaxQueryは区間最大値を扱うSegmentTreeを返します。
func NewRangeMaxQuery() *SegmentTree {
	return &SegmentTree{
		-1,
		[]int{},
		SegmentTreeFunctions{
			func() int { return 0 },
			func(a, b int) int { return max(a, b) },
		},
	}
}

// NewRangeMinQueryは区間最小値を扱うSegmentTreeを返します。
// tested:
//   https://judge.yosupo.jp/problem/staticrmq
func NewRangeMinQuery() *SegmentTree {
	return &SegmentTree{
		-1,
		[]int{},
		SegmentTreeFunctions{
			func() int { return INF18 },
			func(a, b int) int { return min(a, b) },
		},
	}
}

// initは[0, n)のsegment treeを初期化します。
// 各要素の値は単位元となります。
// tested:
//   https://atcoder.jp/contests/abl/tasks/abl_d
func (st *SegmentTree) init(n int) {
	// xはn*2を超える最小の2べき
	x := 1
	for x/2 < n+1 {
		x *= 2
	}
	st.n = x / 2
	st.nodes = make([]int, x)
	for i := 0; i < x; i++ {
		st.nodes[i] = st.f.e()
	}
}

// initAsArrayはvalsで配列を初期化します。
// 区間の長さはlen(vals)になります。
// tested:
//   https://judge.yosupo.jp/problem/staticrmq
func (st *SegmentTree) initAsArray(vals []int) {
	n := len(vals)
	// xはn*2を超える最小の2べき
	x := 1
	for x/2 < n {
		x *= 2
	}
	st.n = x / 2
	st.nodes = make([]int, x)

	for i, v := range vals {
		st.nodes[i+st.n] = v
	}
	for i := st.n - 1; i > 0; i-- {
		st.nodes[i] = st.f.calc(st.nodes[i*2], st.nodes[i*2+1])
	}
}

// queryはi番目の値を取得します。
func (st *SegmentTree) query(i int) int {
	t := i + st.n
	ret := st.nodes[t]

	for {
		t /= 2
		if t == 0 {
			break
		}
		ret = st.f.calc(ret, st.nodes[t])
	}

	return ret
}

// update は[l, r)の値にvalueを適用します。
func (st *SegmentTree) update(l, r, value int) {
	for ll, rr := l+st.n, r+st.n; ll < rr; ll, rr = ll/2, rr/2 {
		if ll%2 == 1 {
			st.nodes[ll] = st.f.calc(st.nodes[ll], value)
			ll++
		}
		if rr%2 == 1 {
			rr--
			st.nodes[rr] = st.f.calc(st.nodes[rr], value)
		}
	}
}

// getはi番目(0-based)の要素を返します。
func (st *SegmentTree) get(i int) int {
	return st.nodes[i+st.n]
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

// compress0はnumbersで渡した値を0から座標圧縮します。
func compress0(numbers map[int]int) (map[int]int, []int) {
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

// compress0はnumbersで渡した値を0から座標圧縮します。
func compress1(numbers map[int]int) (map[int]int, map[int]int) {
	keys := sort.IntSlice{}
	for i := range numbers {
		keys = append(keys, i)
	}
	sort.Sort(keys)
	values := map[int]int{}
	for i, v := range keys {
		numbers[v] = i + 1
		values[i+1] = v
	}

	return numbers, values
}
