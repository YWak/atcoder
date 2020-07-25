package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	A := make([]int, N)

	money := 1000
	kabu := 0

	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	for i := 0; i < N-1; i++ {
		if A[i] < A[i+1] {
			// 翌日上がるなら全力で買う
			c := money / A[i]
			kabu += c
			money -= c * A[i]
		} else {
			// 下がってたら全部売る
			money += kabu * A[i]
			kabu = 0
		}
	}

	for i := N - 2; i >= 0; i-- {
		if A[i] < A[i+1] {
			money += kabu * A[i+1]
			break
		}
	}

	fmt.Println(money)
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
	fmt.Fprintln(os.Stderr, args)
}
