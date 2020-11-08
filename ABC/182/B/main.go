package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	A := nextInts(N)

	counts := make([]int, 1001)

	for i := 0; i < N; i++ {
		for k := 1; k*k <= A[i]; k++ {
			if A[i]%k == 0 {
				counts[k]++

				if A[i]/k != k {
					counts[A[i]/k]++
				}
			}
		}
	}

	ci := 0
	for i := 2; i < len(counts); i++ {
		if counts[ci] <= counts[i] {
			ci = i
		}
	}

	fmt.Println(ci)
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
