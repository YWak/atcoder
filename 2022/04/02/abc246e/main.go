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

type q struct {
	x, y, dir int
}

func calc() {
	n := in.NextInt()
	ax, ay := in.NextInt2d(-1, -1)
	bx, by := in.NextInt2d(-1, -1)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = in.NextString()
	}

	dirs := [][]int{
		{+1, +1},
		{+1, -1},
		{-1, -1},
		{-1, +1},
	}

	cost := make([][][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			cost[i][j] = []int{INF18, INF18, INF18, INF18}
		}
	}
	ok := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < n && s[x][y] != '#'
	}
	queue := NewDeque()

	for d, dir := range dirs {
		cost[ax][ay][d] = 1
		x, y := ax+dir[0], ay+dir[1]
		if ok(x, y) {
			queue.PushFront(&q{ax, ay, d})
		}
	}
	for queue.Len() > 0 {
		item := queue.PopFront()
		// debug(item.x, item.y, item.dir)
		if item.x == bx && item.y == by {
			out.Println(cost[item.x][item.y][item.dir])
			return
		}

		for d1, dir := range dirs {
			x, y := item.x+dir[0], item.y+dir[1]
			if !ok(x, y) {
				continue
			}

			if d1 == item.dir && cost[x][y][d1] > cost[item.x][item.y][item.dir] {
				queue.PushFront(&q{x, y, d1})
				cost[x][y][d1] = cost[item.x][item.y][item.dir]
			} else if cost[x][y][d1] > cost[item.x][item.y][item.dir] {
				queue.PushBack(&q{x, y, d1})
				cost[x][y][d1] = cost[item.x][item.y][item.dir] + 1
			}
		}
	}

	out.Println(-1)
}

// DequeNode はDequeの各要素を保持するstruct
type DequeNode struct {
	value *q
	prev  *DequeNode
	next  *DequeNode
}

// Deque は両端の操作が可能なキューです
type Deque struct {
	head     *DequeNode
	tail     *DequeNode
	length   int
	reversed bool
}

// NewDeque はDequeを作成します
func NewDeque() Deque {
	return Deque{nil, nil, 0, false}
}

// Len はDequeに含まれる要素の数を取得します
func (deque *Deque) Len() int {
	return deque.length
}

// PushFront はDequeの先頭に値を追加します。
func (deque *Deque) PushFront(value *q) {
	if deque.reversed {
		deque.pushBackInternal(value)
	} else {
		deque.pushFrontInternal(value)
	}
}

// PushBack はDequeの末尾に値を追加します。
func (deque *Deque) PushBack(value *q) {
	if deque.reversed {
		deque.pushFrontInternal(value)
	} else {
		deque.pushBackInternal(value)
	}
}

// PopFront はDequeの先頭から値を取得し、値を除去します。
func (deque *Deque) PopFront() *q {
	if deque.reversed {
		return deque.popBackInternal()
	}
	return deque.popFrontInternal()
}

// PopBack はDequeの末尾から値を取得し、値を除去します。
func (deque *Deque) PopBack() *q {
	if deque.reversed {
		return deque.popFrontInternal()
	}
	return deque.popBackInternal()
}

// Front はDequeの先頭の値を取得します。
func (deque *Deque) Front() *q {
	if deque.reversed {
		return deque.tail.value
	}
	return deque.head.value
}

// Back はDequeの末尾の値を取得します。
func (deque *Deque) Back() *q {
	if deque.reversed {
		return deque.head.value
	}
	return deque.tail.value
}

// Reverse はDequeの順序を逆転します。
func (deque *Deque) Reverse() {
	deque.reversed = !deque.reversed
}

// ToArray はDequeを配列化します。
func (deque *Deque) ToArray() []*q {
	ret := make([]*q, 0, deque.Len())

	if deque.reversed {
		for p := deque.tail; p != nil; p = p.prev {
			ret = append(ret, p.value)
		}
	} else {
		for p := deque.head; p != nil; p = p.next {
			ret = append(ret, p.value)
		}
	}
	return ret
}

func (deque *Deque) pushBackInternal(value *q) {
	node := DequeNode{value, deque.tail, nil}
	if deque.tail == nil {
		deque.head = &node
	} else {
		deque.tail.next = &node
	}
	deque.tail = &node
	deque.length++
}

func (deque *Deque) pushFrontInternal(value *q) {
	node := DequeNode{value, nil, deque.head}
	if deque.head == nil {
		deque.tail = &node
	} else {
		deque.head.prev = &node
	}
	deque.head = &node
	deque.length++
}

func (deque *Deque) popFrontInternal() *q {
	node := deque.head
	deque.head = node.next
	if deque.head == nil {
		deque.tail = nil
	} else {
		deque.head.prev = nil
	}

	node.prev = nil
	node.next = nil
	deque.length--
	return node.value
}

func (deque *Deque) popBackInternal() *q {
	node := deque.tail
	deque.tail = node.prev
	if deque.tail == nil {
		deque.head = nil
	} else {
		deque.tail.next = nil
	}

	node.next = nil
	node.prev = nil
	deque.length--

	return node.value
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
