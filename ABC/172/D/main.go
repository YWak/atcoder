package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var counts = make([]int, 10000000+10)

func main() {
	N := nextInt()

	for i := 1; i < len(counts); i++ {
		for j := i; j < len(counts); j += i {
			counts[j]++
		}
	}

	s := int64(0)

	for i := 1; i <= N; i++ {
		s += int64(i) * int64(counts[i])
	}

	fmt.Println(s)
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
