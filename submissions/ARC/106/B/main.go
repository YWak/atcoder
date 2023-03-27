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

	a := make([]int64, N)
	b := make([]int64, N)

	for i := 0; i < N; i++ {
		a[i] = nextInt64()
	}
	for i := 0; i < N; i++ {
		b[i] = nextInt64()
	}

	uf := ufNew(N)

	for i := 0; i < M; i++ {
		c := nextInt() - 1
		d := nextInt() - 1

		uf.Unite(c, d)
	}

	suma := make([]int64, N)
	sumb := make([]int64, N)

	for i := 0; i < N; i++ {
		r := uf.Root(i)
		suma[r] += a[i]
		sumb[r] += b[i]
	}
	for i := 0; i < N; i++ {
		if suma[i] != sumb[i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

// UnionFind : UnionFind構造を保持する構造体
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

// Root はxのルートを得る
func (uf *UnionFind) Root(x int) int {
	p := x
	for p != uf.par[p] {
		p = uf.par[p]
	}
	uf.par[x] = p
	return p
}

// Unite はxとyを併合する。集合の構造が変更された(== 呼び出し前は異なる集合だった)かどうかを返す
func (uf *UnionFind) Unite(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)

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

// Same はxとyが同じノードにいるかを判断する
func (uf *UnionFind) Same(x, y int) bool {
	rx := uf.Root(x)
	ry := uf.Root(y)
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

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
