package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	L := nextInt()

	hurdles := make([]bool, L+10)
	for i := 0; i < N; i++ {
		x := nextInt()
		hurdles[x] = true
	}

	T1 := nextInt()
	T2 := nextInt()
	T3 := nextInt()

	dp := make([]int, L+10)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 0; i <= L; i++ {
		p := sel(hurdles[i], T3, 0)
		// 行動1
		dp[i+1] = min(dp[i+1], dp[i]+T1+p)

		// 行動2
		if i+1 == L {
			dp[i+1] = min(dp[i+1], dp[i]+T1+T2/2+p)
		} else {
			dp[i+2] = min(dp[i+2], dp[i]+T1+T2+p)
		}
		// 行動3
		if i+1 == L {
			dp[L] = min(dp[L], dp[i]+T1/2+T2*1/2+p)
		} else if i+2 == L {
			dp[L] = min(dp[L], dp[i]+T1/2+T2*3/2+p)
		} else if i+3 == L {
			dp[L] = min(dp[L], dp[i]+T1/2+T2*5/2+p)
		} else {
			dp[i+4] = min(dp[i+4], dp[i]+T1+T2*3+p)
		}
	}

	fmt.Println(dp[L])
}

func sel(cond bool, a, b int) int {
	if cond {
		return a
	}
	return b
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
