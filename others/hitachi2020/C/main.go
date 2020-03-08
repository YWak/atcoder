package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var distance []int
var graph [][]int

func main() {
	N := nextInt()

	graph = make([][]int, N)
	distance = make([]int, N)

	// グラフを構成する O(n)
	for i := 0; i < N-1; i++ {
		a := nextInt() - 1
		b := nextInt() - 1

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// 距離を構成する O(n)
	dfs(-1, 0, 0)
	root := 0
	for i := 0; i < N; i++ {
		if distance[root] < distance[i] {
			root = i
		}
	}
	dfs(-1, root, 0)

	// 距離をマップにする O(n)
	dist := map[int][]int{}
	for i := 0; i < N; i++ {
		d := distance[i]
		l := dist[d]

		dist[d] = append(l, i)
	}

	// 順列を作る O(n)
	P := make([]int, N)

	for i := 0; i < N; i++ {
		P[i] = i + 1
	}

	// 順列を構成する
	for i := 0; i < N; i++ {
		//j := i + 3

	}

	// 出力 O(n)
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", P[i])
	}
	fmt.Println()
}

func dfs(prev, current, d int) {
	for i := 0; i < len(graph[current]); i++ {
		c := graph[current][i]

		if c != prev {
			distance[c] = d + 1
			dfs(current, c, d+1)
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
