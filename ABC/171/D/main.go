package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	count := map[int64]int64{}

	sum := int64(0)
	for i := 0; i < N; i++ {
		a := nextInt64()
		n, ok := count[a]

		sum += a
		if !ok {
			count[a] = 1
		} else {
			count[a] = n + 1
		}
	}

	Q := nextInt()
	for i := 0; i < Q; i++ {
		b := nextInt64()
		c := nextInt64()

		n, ok := count[b]

		if !ok {
			n = 0
		}
		count[b] = 0

		m, ok := count[c]

		if !ok {
			m = 0
		}
		count[c] = n + m

		sum = sum - b*n + c*(n)
		fmt.Println(sum)
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

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args)
}
