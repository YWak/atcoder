package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	X := abs(nextInt64())
	K := nextInt64()
	D := nextInt64()

	n := X / D

	if n >= K {
		fmt.Println(X - D*K)
	} else if (K-n)%2 == 0 {
		fmt.Println(abs(X - D*n))
	} else {
		fmt.Println(abs(X - D*(n+1)))
	}
}

func abs(a int64) int64 {
	if a > 0 {
		return a
	}
	return -a
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
