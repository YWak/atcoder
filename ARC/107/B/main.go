package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	N := nextInt()
	K := nextInt()

	c := 0

	for ab := 2; ab <= 2*N; ab++ {
		cd := ab - K
		c += count(N, ab) * count(N, cd)
	}

	fmt.Println(c)
}

func count(N, a int) int {
	if a > 2*N || a < 2 {
		return 0
	}
	if a <= N+1 {
		return a - 1
	}

	return N*2 + 1 - a
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
