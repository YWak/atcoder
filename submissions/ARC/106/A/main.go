package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt64()

	c := int64(1)
	MAX := int64(1e18)

loop:
	for a := 1; a < 39; a++ {
		c *= 3

		if c >= MAX || c >= N {
			break
		}

		n := N - c
		b := 0
		for n > 1 {
			if n%5 != 0 {
				continue loop
			}

			n = n / 5
			b++
		}

		if b == 0 {
			continue
		}
		fmt.Printf("%d %d\n", a, b)
		return
	}

	fmt.Println(-1)
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
