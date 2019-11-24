package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var graph [][]edge
var colors []int
var n int

func main() {
	n = nextInt()
	graph = make([][]edge, n+1)
	colors = make([]int, n)
	c := 0
	graph[0] = []edge{edge{node: 1, index: 0}}

	for i := 1; i < n; i++ {
		a := nextInt()
		b := nextInt()
		ea := edge{node: b, index: i}
		eb := edge{node: a, index: i}
		graph[a] = append(graph[a], ea)
		graph[b] = append(graph[b], eb)

		c = max(c, len(graph[a]))
		c = max(c, len(graph[b]))
	}

	// 適当に塗っていく
	colors[0] = 0
	dfs(0, 0)

	fmt.Println(c)
	for i := 1; i < n; i++ {
		fmt.Println(colors[i])
	}
}

type edge struct {
	node  int
	index int
}

func dfs(prev, curr int) {
	var prevColor int
	nextColor := 1
	nexts := graph[curr]

	for i := 0; i < len(nexts); i++ {
		next := nexts[i]
		if next.node == prev {
			prevColor = colors[next.index]
			break
		}
	}

	for i := 0; i < len(nexts); i++ {
		next := nexts[i]
		if next.node == prev {
			continue
		}
		if nextColor == prevColor {
			nextColor++
		}

		colors[next.index] = nextColor
		dfs(curr, next.node)
		nextColor++
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
