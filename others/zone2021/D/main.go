//lint:file-ignore U1000 using template
package main

import (
	"bufio"
	"fmt"
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

func main() {
	fmt.Println(solve(nextString()))
	// random()
}

func random() {
	c := "qwertyuiopasdfghjklzxcvbnmR"

	for i := 0; i < 1000; i++ {
		n := rand.Intn(500)
		b := make([]byte, n)
		for j := 0; j < n; j++ {
			b[j] = c[rand.Intn(len(c))]
		}
		s := string(b)
		s1, s2 := solve(s), naive(s)
		if s1 == s2 {
			// fmt.Println("OK", s)
		} else {
			fmt.Println("NG", s, s1, s2)
		}
	}
}

func naive(s string) string {
	b := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			for k, l := 0, len(b)-1; k < l; k, l = k+1, l-1 {
				b[k], b[l] = b[l], b[k]
			}
		} else if len(b) > 0 && b[len(b)-1] == s[i] {
			b = b[:len(b)-1]
		} else {
			b = append(b, s[i])
		}
	}

	return string(b)
}

func solve(s string) string {
	deque := NewDeque()

	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			deque.Reverse()
		} else if deque.Len() > 0 && deque.PeekBack() == int(s[i]) {
			deque.PopBack()
		} else {
			deque.PushBack(int(s[i]))
		}
	}
	array := deque.ToArray()
	arr := make([]byte, 0, len(array))

	for i := 0; i < len(array); i++ {
		arr = append(arr, byte(array[i]))
	}
	if deque.Len() != len(array) {
		debug(fmt.Sprintf("%d vs %d", deque.length, len(array)))
	}

	return string(arr)
}

func debugrev(hh, ht, th, tt []byte) {
	b1 := make([]byte, len(hh))
	b2 := make([]byte, len(th))
	for i := 0; i < len(hh); i++ {
		b1[i] = hh[len(hh)-1-i]
	}
	for i := 0; i < len(th); i++ {
		b2[i] = th[len(th)-1-i]
	}
	debug(string(b1), string(ht), string(b2), string(tt))
}
func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

// ==================================================
// 入力操作
// ==================================================
// DequeNode はDequeの各要素を保持するstruct
type DequeNode struct {
	value int
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
func (deque *Deque) PushFront(value int) {
	if deque.reversed {
		_dequePushBackInternal(deque, value)
	} else {
		_dequePushFrontInternal(deque, value)
	}
}

// PushBack はDequeの末尾に値を追加します。
func (deque *Deque) PushBack(value int) {
	if deque.reversed {
		_dequePushFrontInternal(deque, value)
	} else {
		_dequePushBackInternal(deque, value)
	}
}

// PopFront はDequeの先頭から値を取得し、値を除去します。
func (deque *Deque) PopFront() int {
	if deque.reversed {
		return _dequePopBackInternal(deque)
	}
	return _dequePopFrontInternal(deque)
}

// PopBack はDequeの末尾から値を取得し、値を除去します。
func (deque *Deque) PopBack() int {
	if deque.reversed {
		return _dequePopFrontInternal(deque)
	}
	return _dequePopBackInternal(deque)
}

// PeekFront はDequeの先頭の値を取得します。
func (deque *Deque) PeekFront() int {
	if deque.reversed {
		return deque.tail.value
	}
	return deque.head.value
}

// PeekBack はDequeの末尾の値を取得します。
func (deque *Deque) PeekBack() int {
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
func (deque *Deque) ToArray() []int {
	ret := make([]int, 0, deque.Len())

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

func _dequePushBackInternal(deque *Deque, value int) {
	node := DequeNode{value, deque.tail, nil}
	if deque.tail == nil {
		deque.head = &node
	} else {
		deque.tail.next = &node
	}
	deque.tail = &node
	deque.length++
}

func _dequePushFrontInternal(deque *Deque, value int) {
	node := DequeNode{value, nil, deque.head}
	if deque.head == nil {
		deque.tail = &node
	} else {
		deque.head.prev = &node
	}
	deque.head = &node
	deque.length++
}

func _dequePopFrontInternal(deque *Deque) int {
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

func _dequePopBackInternal(deque *Deque) int {
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
