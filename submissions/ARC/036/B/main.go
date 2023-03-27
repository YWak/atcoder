package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	H := make([]int, N)
	s := make([]int, N)
	u := make([]int, N)

	for i := 0; i < N; i++ {
		H[i] = nextInt()
	}
	s[0] = 1
	u[N-1] = N
	for t := 1; t < N; t++ {
		if H[t-1] < H[t] {
			s[t] = s[t-1]
		} else {
			s[t] = t + 1
		}
	}
	for t := N - 2; t >= 0; t-- {
		if H[t] > H[t+1] {
			u[t] = u[t+1]
		} else {
			u[t] = t + 1
		}
	}

	// debug(s)
	// debug(u)
	c := 1
	for t := 0; t < N; t++ {
		c = max(c, u[t]-s[t]+1)
	}

	fmt.Println(c)
}

func max(a, b int) int {
	if a > b {
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

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
