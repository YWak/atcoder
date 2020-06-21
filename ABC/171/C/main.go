package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var alpha = "abcdefghijklmnopqrstuvwxyz"

func main() {
	N := nextInt64() - 1
	name := ""

	for N >= 0 {
		n := N % 26

		name = fmt.Sprintf("%c%s", alpha[n], name)
		N = N/26 - 1
	}

	fmt.Println(name)
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
