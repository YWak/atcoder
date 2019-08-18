package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	Q := nextInt()

	graph := make([][]int, N+1)
	parents := make([]int, N+1)
	query := make([]int64, N+1)
	ans := make([]int64, N+1)

	for i := 0; i < N-1; i++ {
		a := nextInt()
		b := nextInt()

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// 親を決める
	queue := list.New()
	queue.PushBack(1)
	for queue.Len() > 0 {
		e := queue.Front()
		s := e.Value.(int)
		queue.Remove(e)

		for i := 0; i < len(graph[s]); i++ {
			if graph[s][i] == parents[s] {
				continue
			}
			n := graph[s][i]
			parents[n] = s
			queue.PushBack(n)
		}
	}

	for i := 0; i < Q; i++ {
		p := nextInt()
		x := int64(nextInt())
		query[p] += x
	}

	// fmt.Println(query)
	queue.PushBack(1)
	ans[1] = query[1]
	for queue.Len() > 0 {
		e := queue.Front()
		s := e.Value.(int)
		queue.Remove(e)

		for i := 0; i < len(graph[s]); i++ {
			if graph[s][i] == parents[s] {
				continue
			}
			n := graph[s][i]
			ans[n] += query[n] + ans[parents[n]]
			queue.PushBack(n)
		}
	}
	for i := 1; i <= N; i++ {
		fmt.Print(ans[i])

		if i != N {
			fmt.Print(" ")
		}
	}
	fmt.Println()
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
