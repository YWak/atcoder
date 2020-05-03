package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	M := nextInt()
	H := make([]int, N)
	root := make([]int, N)

	for i := 0; i < N; i++ {
		H[i] = nextInt()
		root[i] = i
	}

	for i := 0; i < M; i++ {
		a := nextInt() - 1
		b := nextInt() - 1

		// H[root[a]]とH[root[b]]を比べる
		if root[a] == root[b] {
			continue
		}
		if H[root[a]] > H[root[b]] {
			root[b] = root[a]
		} else {
			root[a] = root[b]
		}
	}
	c := 0
	hist := make([]bool, N)

	for i := 0; i < N; i++ {
		if !hist[root[i]] {
			c++
			hist[root[i]] = true
		}
	}
	fmt.Println(c)
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
