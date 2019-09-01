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

	for i := 0; i < N; i++ {
		H[i] = nextInt()
	}
	m := 0
	s[N-1] = 0
	for i := 1; i < N; i++ {
		n := N - i - 1

		if H[n] < H[n+1] {
			s[n] = 0
		} else {
			s[n] = s[n+1] + 1
		}
		m = max(s[n], m)
	}

	fmt.Println(m)
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
