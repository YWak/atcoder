package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	M := nextInt()
	P := make([]int, N)
	for i := 0; i < N; i++ {
		P[i] = nextInt()
	}

	uf := ufNew(N)

	for i := 0; i < M; i++ {
		x := nextInt()
		y := nextInt()

		ufUnite(&uf, x-1, y-1)
	}

	c := 0
	for i := 0; i < N; i++ {
		if ufSame(&uf, i, P[i]-1) {
			c++
		}
	}

	fmt.Println(c)
}

// UnionFind : 森構造を保持する構造体
type UnionFind struct {
	par  []int // i番目のノードに対応する親
	rank []int // i番目のノードの階層
}

// [0, n)のノードを持つUnion-Findを作る
func ufNew(n int) UnionFind {
	uf := UnionFind{par: make([]int, n), rank: make([]int, n)}

	for i := 0; i < n; i++ {
		uf.par[i] = i
	}

	return uf
}

// xのルートを得る
func ufRoot(uf *UnionFind, x int) int {
	p := x
	for p != uf.par[p] {
		p = uf.par[p]
	}
	uf.par[x] = p
	return p
}

// xとyを併合する。集合の構造が変更された(== 呼び出し前は異なる集合だった)かどうかを返す
func ufUnite(uf *UnionFind, x, y int) bool {
	rx := ufRoot(uf, x)
	ry := ufRoot(uf, y)

	if rx == ry {
		return false
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
	uf.par[ry] = rx
	return true
}

// xとyが同じノードにいるかを判断する
func ufSame(uf *UnionFind, x, y int) bool {
	rx := ufRoot(uf, x)
	ry := ufRoot(uf, y)
	return rx == ry
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

func nextBytes() []byte {
	stdin.Scan()
	return stdin.Bytes()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func nextInt64() int64 {
	i, _ := strconv.ParseInt(nextString(), 10, 64)
	return i
}
