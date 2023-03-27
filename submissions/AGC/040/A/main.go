package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S := nextBytes()
	N := len(S) + 1
	c := int64(0)

	C := make([]byte, 1)
	L := make([]int, 1)
	n := 0

	C[0] = S[0]
	L[0] = 1

	for i := 1; i < N-1; i++ {
		if C[n] != S[i] {
			C = append(C, S[i])
			L = append(L, 0)
			n++
		}
		L[n]++
	}
	if C[n] == '<' {
		C = append(C, '>')
		L = append(L, 0)
	}

	n = 0
	if C[0] == '>' {
		c = c + sum(L[0])
		n++
	}
	for ; n < len(C); n += 2 {
		if L[n] < L[n+1] {
			c = c + sum(L[n]-1) + sum(L[n+1])
		} else {
			c = c + sum(L[n]) + sum(L[n+1]-1)
		}
	}

	fmt.Println(c)
}

func sum(n int) int64 {
	return int64((n * (n + 1)) / 2)
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
