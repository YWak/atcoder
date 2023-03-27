package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A := nextInt()
	B := nextInt()

	fmt.Println(max(A+B, A-B, A*B))
}

func max(a int, b ...int) int {
	l := len(b)
	for i := 0; i < l; i++ {
		if a < b[i] {
			a = b[i]
		}
	}

	return a
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
