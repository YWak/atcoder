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

	A := make([]int, M)
	C := make([]int, M)

	for i := 0; i < M; i++ {
		A[i] = nextInt()
		b := nextInt()

		for j := 0; j < b; j++ {
			c := nextInt() - 1
			C[i] |= (1 << uint(c))
		}
	}
	inf := 1 << 29

	max := 1 << uint(N)
	dp := make([]int, max)

	for i := 0; i < max; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	for s := 0; s < max; s++ {
		for i := 0; i < M; i++ {
			t := s | C[i]
			cost := dp[s] + A[i]
			dp[t] = min(dp[t], cost)
		}
	}

	ans := dp[max-1]
	if ans == inf {
		ans = -1
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
