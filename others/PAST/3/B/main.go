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
	Q := nextInt()

	score := make([]int, M)
	ans := make([][]bool, N)

	for i := 0; i < N; i++ {
		ans[i] = make([]bool, M)
	}

	for i := 0; i < Q; i++ {
		s := nextInt()
		n := nextInt() - 1
		if s == 1 {
			c := 0

			for j := 0; j < M; j++ {
				if ans[n][j] {
					c += N - score[j]
				}
			}
			fmt.Println(c)
		} else {
			m := nextInt() - 1
			ans[n][m] = true
			score[m]++
		}
	}
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
