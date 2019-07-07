package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	L := nextInt()
	R := nextInt()
	mod := 2019

	from := L % mod
	len := min(R, L+mod-1) - L + 1
	m := mod

	for i := 0; i < len-1; i++ {
		l := (from + i) % mod
		for j := i + 1; j < len; j++ {
			r := (from + j) % mod
			v := (l * r) % mod

			if v < m {
				m = v
			}
			if v == 0 {
				break
			}
		}
	}

	fmt.Println(m)
}

var stdin = initStdin()

func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}

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
