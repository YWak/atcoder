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
	C := nextInt()
	D := nextInt()

	cookies := []int{A, B, C, D}
	all := A + B + C + D

	for i := 1; i <= 15; i++ {
		c := 0
		for j := 0; j < 4; j++ {
			if (i>>j)&1 == 1 {
				c += cookies[j]
			}
			if all-c == c {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
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
