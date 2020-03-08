package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var a []int64
var b []int64

func main() {

	N := nextInt()
	T := nextInt64()

	a = make([]int64, N)
	b = make([]int64, N)

	for i := 0; i < N; i++ {
		a[i] = nextInt64()
		b[i] = nextInt64()
	}

	fmt.Println()
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
