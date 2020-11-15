package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var H int
var W int
var S [][]byte

func solve() int {
	dp := make([][]int, H)
	x := make([][]int, H)
	y := make([][]int, H)
	z := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
		x[i] = make([]int, W)
		y[i] = make([]int, W)
		z[i] = make([]int, W)
	}

	dp[0][0] = 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if S[i][j] == '#' {
				continue
			}
			if j > 0 {
				x[i][j] = (x[i][j-1] + dp[i][j-1]) % mod
			}
			if i > 0 {
				y[i][j] = (y[i-1][j] + dp[i-1][j]) % mod
			}
			if i > 0 && j > 0 {
				z[i][j] = (z[i-1][j-1] + dp[i-1][j-1]) % mod
			}
			dp[i][j] = (x[i][j] + y[i][j] + z[i][j]) % mod
		}
	}
	return dp[H-1][W-1]
}

func main() {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	H = nextInt()
	W = nextInt()

	S = make([][]byte, H)
	for i := 0; i < H; i++ {
		S[i] = nextBytes()

		if len(S[i]) != W {
			debug("NG!!", len(S[i]), W)
		}
	}

	fmt.Fprintln(wr, solve())
}

type mint int

const mod = int(1e9 + 7)

var stdin = initStdin()

func initStdin() *bufio.Scanner {
	bufsize := 1 * 1024 * 1024 // 1 MB
	var stdin = bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, bufsize), bufsize)
	stdin.Split(bufio.ScanWords)
	return stdin
}

// NextString は標準入力から文字列を読み込みます。
func NextString() string {
	stdin.Scan()
	return stdin.Text()
}

func nextBytes() []byte {
	return []byte(NextString())
}

func nextInt() int {
	i, _ := strconv.Atoi(NextString())
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
