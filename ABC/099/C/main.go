package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	a := 1
	b := 1

	c := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if a*6 <= i {
			a *= 6
		}
		if b*9 <= i {
			b *= 9
		}

		c[i] = c[i-1] + 1

		for j := 1; j < 6 && j*a <= i; j++ {
			c[i] = min(c[i], c[i-j*a]+j)
		}
		for j := 1; j < 9 && j*b <= i; j++ {
			c[i] = min(c[i], c[i-j*b]+j)
		}
	}

	// debug(c)
	fmt.Println(c[N])
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

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
