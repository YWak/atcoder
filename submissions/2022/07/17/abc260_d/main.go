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
const INF18 = int(2e18) + int(2e9)

// INF9 は最大値を表す数
const INF9 = int(2e9)

// N10_6は10^6
const N10_6 = int(1e6)

var in *In
var out *Out

type Treap struct {
	root *node
	// allowDup は、このTreapのキーとして重複を許すかどうか
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

// NewIntTreap は、int型のキーを使用し、昇順に保存するTreapを作成して返します。
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

// NewRevIntTreap は、int型のキーを使用し、降順に保存するTreapを作成して返します。
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

// Len は、このTreapに含まれる要素の数を返します。
func (t *Treap) Len() int {
	return t._count(t.root)
}

// Get は、keyに対応する値を探して返します。 存在しない場合はnilを返します。
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

// GetKthは k (1-indexed)番目のキーと対応する値を返します。
// 該当する要素が存在しない場合は (nil, nil) を返します。
func (t *Treap) GetKth(k int) (interface{}, interface{}) {
	if k < 0 || t.Len() < k {
		return nil, nil
	}
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

// Find は、key以下で最大のキーを返します。
// 対応するキーがなければnilを返します。
func (t *Treap) Find(key interface{}) interface{} {
	n := t.root
	var ans *node

	for n != nil {
		c := t.comparator(key, n.key)
		if c == 0 {
			return key
		}
		if c < 0 {
			n = n.left
		} else {
			ans = n
			n = n.right
		}
	}

	if ans == nil {
		return nil
	}
	return ans.key
}

// Put は、keyとそれに対応するvalueを保存し、古い値を返します。
// すでにキーが登録されている場合、allowDupがtrueなら挿入、falseなら上書きされます。
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
	if c <= 0 {
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

func calc() {
	n, k := in.NextInt2()
	p := in.NextInts(n)
	on := make([]int, n+1) // 何番の上にのっているか
	for i := 0; i < n+1; i++ {
		on[i] = -1
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	m := NewRevIntTreap()

	for i, x := range p {
		t := m.Find(x)
		if t == nil {
			if k > 1 {
				m.Put(x, 1)
				// debug("put first", x)
			} else {
				ans[x-1] = i + 1
			}
			continue
		}
		c := m.Get(t).(int)
		c++
		on[x] = t.(int)
		// debug("put", x, "on", t)
		if c == k {
			// 全削除
			for q := x; q != -1; q = on[q] {
				// debug("remove", q, "at", i+1)
				ans[q-1] = i + 1
			}
		} else {
			// 追加
			m.Put(x, c)
		}
		m.Remove(t)
	}
	for _, v := range ans {
		out.Println(v)
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

// chmax は aとbのうち大きい方をaに設定します。
func chmax(a *int, b int) {
	*a = max(*a, b)
}

// chmin は aとbのうち小さい方をaに設定します。
func chmin(a *int, b int) {
	*a = min(*a, b)
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

// chiはcondがtrueのときok, falseのときngを返します。
func chi(cond bool, ok, ng int) int {
	if cond {
		return ok
	}
	return ng
}

// chbはcondがtrueのときok, falseのときngを返します。
func chb(cond bool, ok, ng byte) byte {
	if cond {
		return ok
	}
	return ng
}

// chsはcondがtrueのときok, falseのときngを返します。
func chs(cond bool, ok, ng string) string {
	if cond {
		return ok
	}
	return ng
}

// extmulはa*bの結果を返します。
// 2つ目の値が+1ならオーバーフロー、-1ならアンダーフローが発生したことを表します。
func extmul(a, b int) (int, int) {
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
