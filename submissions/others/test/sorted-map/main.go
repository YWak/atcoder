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

type ll int
type lls []ll

func (a ll) Compare(bb SortedMapKey) int {
	b := bb.(ll)
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func calc() {
	m := NewSortedMap()

	for i := 0; i < 15; i++ {
		k := ll(rand.Int())
		v := ll(rand.Int())
		m.Put(k, v)
	}
	hasNext, next := m.Iterate()
	for hasNext() {
		k, v := next()
		fmt.Printf("% 20d % 20d\n", k, v)
	}

	fmt.Println()
}

func main() {
	calc()
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

// Randmoized Binary Search Treeの実装
// https://tjkendev.github.io/procon-library/python/binary_search_tree/RBST.html
type SortedMapKey interface {
	Compare(b SortedMapKey) int
}

type SortedMapNode struct {
	key   SortedMapKey
	value interface{}
	size  int
	left  *SortedMapNode
	right *SortedMapNode
}

type SortedMap struct {
	root *SortedMapNode
	size int
}

func NewSortedMap() SortedMap {
	return SortedMap{nil, 0}
}

// Put は keyに対応する値を設定します。
func (m *SortedMap) Put(key SortedMapKey, value interface{}) {
	pp := &m.root
	p := m.root

	for p != nil {
		c := key.Compare(p.key)
		if c == 0 {
			p.value = value
			return
		}
		if c < 0 {
			pp = &p.left
			p = p.left
		} else {
			pp = &p.right
			p = p.right
		}
	}
	*pp = &SortedMapNode{key, value, 0, nil, nil}
	m.size++
}

// Get は keyに対応する値を取得します。
func (m *SortedMap) Get(key SortedMapKey) (interface{}, bool) {
	p := m.root
	for p != nil {
		c := p.key.Compare(key)
		if c == 0 {
			return p.value, true
		}
		if c < 0 {
			p = p.left
		} else {
			p = p.right
		}
	}

	return nil, false
}

// At は k番目のキーを取得します。
func (m *SortedMap) At(k int) SortedMapKey {
	return nil
}

// Remove は key に対応する値を削除します。
func (m *SortedMap) Remove(key SortedMapKey) interface{} {
	return nil
}

// Iterate は keyが小さい順に値を取り出す関数を返します。
func (m *SortedMap) Iterate() (func() bool, func() (SortedMapKey, interface{})) {
	keys := make([]SortedMapKey, m.size)
	values := make([]interface{}, m.size)

	// 探索
	m.dfs(-1, m.root, &keys, &values)

	k := 0
	next := func() bool {
		return k < len(keys)
	}
	get := func() (SortedMapKey, interface{}) {
		key := keys[k]
		value := values[k]
		k++
		return key, value
	}

	return next, get
}

func (m *SortedMap) dfs(i int, node *SortedMapNode, keys *[]SortedMapKey, values *[]interface{}) int {
	if node == nil {
		return i
	}
	i = m.dfs(i, node.left, keys, values) // 行きがけ
	i++
	(*keys)[i] = node.key
	(*values)[i] = node.value
	return m.dfs(i, node.right, keys, values)
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
