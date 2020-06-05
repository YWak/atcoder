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
	Q := nextInt()

	graph := make([][]int, N)

	for i := 0; i < M; i++ {
		u := nextInt() - 1
		v := nextInt() - 1

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	colors := make([]int, N)
	for i := 0; i < N; i++ {
		c := nextInt()
		colors[i] = c
	}

	for i := 0; i < Q; i++ {
		s := nextInt()

		x := nextInt() - 1
		fmt.Println(colors[x])

		if s == 1 {
			for j := 0; j < len(graph[x]); j++ {
				n := graph[x][j]
				colors[n] = colors[x]
			}
		} else {
			y := nextInt()
			colors[x] = y
		}
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
