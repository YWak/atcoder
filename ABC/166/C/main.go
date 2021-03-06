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
	graph := make([][]int, N)

	for i := 0; i < N; i++ {
		H[i] = nextInt()
	}

	for i := 0; i < M; i++ {
		a := nextInt() - 1
		b := nextInt() - 1

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	c := 0

	for i := 0; i < N; i++ {
		ok := true
		for j := 0; j < len(graph[i]); j++ {
			k := graph[i][j]
			if H[i] <= H[k] {
				ok = false
				break
			}
		}
		if ok {
			c++
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
