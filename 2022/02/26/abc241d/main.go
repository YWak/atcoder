//lint:file-ignore U1000 using template
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/bits"
	"math/rand"
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
	Q := in.NextInt()
	m := NewIntTreap()
	m.allowDup = true

	for q := 0; q < Q; q++ {
		t := in.NextInt()
		x := in.NextInt()

		if t == 1 {
			m.Put(x, 1)
			continue
		}

		var kth int
		k := in.NextInt()
		k--
		if t == 2 {
			// 二分探索でx以下の要素が何番目か求める
			ok, ng := 0, m.Len()+1 // okはxが何番目か
			for abs(ok-ng) > 1 {
				mid := (ok + ng) / 2
				v, _ := m.GetKth(mid)
				if v.(int) <= x {
					ok = mid
				} else {
					ng = mid
				}
			}
			kth = ok - k
		} else {
			// 二分探索でx以上の要素が何番目か求める
			ok, ng := m.Len()+1, 0 // okはxが何番目か
			for abs(ok-ng) > 1 {
				mid := (ok + ng) / 2
				v, _ := m.GetKth(mid)
				if v.(int) >= x {
					ok = mid
				} else {
					ng = mid
				}
			}
			kth = ok + k
		}
		v, _ := m.GetKth(kth)
		// debug("t =", t, "x =", x, "k =", k+1, "m =", kth-(2*t-5)*k, "kth =", kth, v)
		// for i := 1; i <= m.Len(); i++ {
		// 	v, _ := m.GetKth(i)
		// 	debug(i, v)
		// }

		if kth <= 0 || kth > m.Len() || v == nil {
			out.Println(-1)
		} else {
			out.Println(v)
		}
	}
}

type Treap struct {
	root       *node
	allowDup   bool
	comparator func(a, b interface{}) int
}

type node struct {
	key   interface{}
	value interface{}
	pri   int
	cnt   int
	left  *node
	right *node
}

func NewTreap(comparator func(a, b interface{}) int) *Treap {
	return &Treap{nil, false, comparator}
}

func NewIntTreap() *Treap {
	return NewTreap(func(a, b interface{}) int {
		aa, bb := a.(int), b.(int)
		if aa == bb {
			return 0
		}
		if aa < bb {
			return -1
		}
		return +1
	})
}

func NewRevIntTreap() *Treap {
	return NewTreap(func(a, b interface{}) int {
		aa, bb := a.(int), b.(int)
		if aa == bb {
			return 0
		}
		if aa < bb {
			return +1
		}
		return -1
	})
}

func (t *Treap) Len() int {
	return t._count(t.root)
}

func (t *Treap) Get(key interface{}) interface{} {
	n := t.root
	for n != nil {
		c := t.comparator(key, n.key)
		if c == 0 {
			return n.value
		} else if c < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	return nil
}

func (t *Treap) GetKth(k int) (interface{}, interface{}) {
	a, b := t._split(t.root, k)

	n := a
	var p *node = nil
	for n != nil {
		p = n
		n = n.right
	}
	if p == nil {
		return nil, nil
	}
	t.root = t._merge(a, b)

	return p.key, p.value
}

func (t *Treap) Put(key interface{}, value interface{}) interface{} {
	n, v := t._put(t.root, key, value, rand.Intn(1<<60))
	t.root = n
	t._update(t.root)

	return v
}

func (t *Treap) _put(n *node, key, value interface{}, pri int) (*node, interface{}) {
	if n == nil {
		return &node{key, value, pri, 1, nil, nil}, nil
	}
	c := t.comparator(key, n.key)
	if c == 0 && !t.allowDup {
		v := n.value
		n.value = value
		return n, v
	}
	if c < 0 || c == 0 && t.allowDup {
		nn, v := t._put(n.left, key, value, pri)
		n.left = nn
		if n.left.pri < n.pri {
			n = t._rotatel(n)
		}

		return t._update(n), v
	} else {
		nn, v := t._put(n.right, key, value, pri)
		n.right = nn
		if n.right.pri < n.pri {
			n = t._rotater(n)
		}
		return t._update(n), v
	}
}

func (t *Treap) _rotatel(n *node) *node {
	result := n.left
	x := result.right
	result.right = n
	n.left = x
	t._update(n)
	t._update(result)
	return result
}

func (t *Treap) _rotater(n *node) *node {
	result := n.right
	x := result.left
	result.left = n
	n.right = x
	t._update(n)
	t._update(result)
	return result
}

func (t *Treap) Remove(key interface{}) interface{} {
	n, v := t._remove(t.root, key)
	t.root = t._update(n)
	return v
}

func (t *Treap) _remove(n *node, key interface{}) (*node, interface{}) {
	if n == nil {
		return nil, nil
	}
	c := t.comparator(key, n.key)
	if c == 0 {
		// このノードを削除する
		v := n.value
		n = t._merge(n.left, n.right)
		t._update(n)
		return n, v
	}
	if c < 0 {
		r := n
		nn, v := t._remove(n.left, key)
		r.left = nn
		if r.left != nil && r.left.pri < r.pri {
			r = t._rotatel(r)
		}
		return t._update(r), v
	} else {
		r := n
		nn, v := t._remove(n.right, key)
		r.right = nn
		if r.right != nil && r.right.pri < r.pri {
			r = t._rotater(r)
		}
		return t._update(r), v
	}
}

func (t *Treap) _count(n *node) int {
	if n == nil {
		return 0
	}
	return n.cnt
}

func (t *Treap) _update(n *node) *node {
	if n != nil {
		n.cnt = t._count(n.left) + t._count(n.right) + 1
	}

	return n
}

func (t *Treap) _merge(l, r *node) *node {
	if l == nil || r == nil {
		if r == nil {
			return l
		} else {
			return r
		}
	}
	if l.pri < r.pri {
		l.right = t._merge(l.right, r)
		return t._update(l)
	} else {
		r.left = t._merge(l, r.left)
		return t._update(r)
	}
}

func (t *Treap) _split(n *node, nth int) (*node, *node) {
	if n == nil {
		return nil, nil
	}
	if nth <= t._count(n.left) {
		a, b := t._split(n.left, nth)
		n.left = b
		return a, t._update(n)
	} else {
		a, b := t._split(n.right, nth-t._count(n.left)-1)
		n.right = a
		return t._update(n), b
	}
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
