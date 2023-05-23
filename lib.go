//lint:file-ignore U1000 using template
package main

import (
	"math"
	"math/bits"
	"sort"
	"strings"

	mymath "github.com/ywak/atcoder/lib/math"
)

// ==================================================
// 数値操作
// ==================================================

// max は aとbのうち大きい方を返します。
func max(a, b int) int {
	return mymath.Max(a, b)
}

// min は aとbのうち小さい方を返します。
func min(a, b int) int {
	return mymath.Min(a, b)
}

// chmax は aとbのうち大きい方をaに設定します。
func chmax(a *int, b int) {
	mymath.Chmax(a, b)
}

// chmin は aとbのうち小さい方をaに設定します。
func chmin(a *int, b int) {
	mymath.Chmin(a, b)
}

// abs は aの絶対値を返します。
func abs(a int) int {
	return mymath.Abs(a)
}

// pow は aのb乗を返します。
func pow(a, b int) int {
	return mymath.Pow(a, b)
}

// divceil は a/b の結果を正の無限大に近づけるように丸めて返します。
func divceil(a, b int) int {
	return mymath.Divceil(a, b)
}

// divfloor は a/b の結果を負の無限大に近づけるように丸めて返します。
func divfloor(a, b int) int {
	return mymath.Divfloor(a, b)
}

// powmod は (x^n) mod m を返します。
func powmod(x, n, m int) int {
	return mymath.Powmod(x, n, m)
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
