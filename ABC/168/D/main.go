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

	graph := make([][]int, N)

	for i := 0; i < M; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// 幅優先探索で。
	shortests := make([]pair, N)
	for i := 0; i < N; i++ {
		shortests[i] = pair{-1, N + 10}
	}
	shortests[0].size = 0

	done := make([]bool, N)
	queue := make([]int, 0, N)
	queue = append(queue, 0)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if done[p] {
			continue
		}
		done[p] = true

		for i := 0; i < len(graph[p]); i++ {
			n := graph[p][i]
			queue = append(queue, n)

			s := shortests[p].size + 1
			if shortests[n].size > s {
				shortests[n].node = p
				shortests[n].size = s
			}
		}
	}

	fmt.Println("Yes")
	for i := 1; i < N; i++ {
		fmt.Println(shortests[i].node + 1)
	}
}

type pair struct {
	node int
	size int
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
