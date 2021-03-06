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
	K := nextInt64()
	sa := make([]int64, N+1)
	sb := make([]int64, M+1)

	for i := 0; i < N; i++ {
		a := nextInt64()
		sa[i+1] = sa[i] + a
	}
	for i := 0; i < M; i++ {
		b := nextInt64()
		sb[i+1] = sb[i] + b
	}

	maxk := 0

	for i := 0; i < len(sa); i++ {
		ok := -1
		ng := len(sb)

		for abs(ok-ng) > 1 {
			mid := (ok + ng) / 2

			if sa[i]+sb[mid] <= K {
				ok = mid
			} else {
				ng = mid
			}
		}

		if ok >= 0 && sa[i]+sb[ok] <= K {
			maxk = max(maxk, i+ok)
		}
	}

	fmt.Println(maxk)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
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
