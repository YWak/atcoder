package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var N int
var K int
var T [][]int

func main() {
	N = nextInt()
	K = nextInt()

	T = make([][]int, N)
	for i := 0; i < N; i++ {
		T[i] = nextInts(N)
	}

	fmt.Println(dfs(0, 0, map[int]bool{0: true}))
}

func dfs(now, cost int, hist map[int]bool) int {
	if len(hist) == N {
		if cost+T[now][0] == K {
			return 1
		}
		return 0
	}

	c := 0
	for i := 1; i < N; i++ {
		_, ok := hist[i]

		if !ok {
			hist[i] = true
			c += dfs(i, cost+T[now][i], hist)
			delete(hist, i)
		}
	}

	return c
}

func nthbit(a int, n int) int { return int((a >> uint(n)) & 1) }

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

func nextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return a
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
