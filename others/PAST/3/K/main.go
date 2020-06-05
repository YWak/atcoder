package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	Q := nextInt()

	heads := make([]node, N)
	tails := make([]*node, N)
	refs := make([]*node, N)

	for i := 0; i < N; i++ {
		heads[i] = node{-i, nil, nil}
		n := node{i, &heads[i], nil}
		heads[i].next = &n
		tails[i] = &n
		refs[i] = &n
	}

	for i := 0; i < Q; i++ {
		f := nextInt() - 1
		t := nextInt() - 1
		x := nextInt() - 1

		// コンテナ
		n := refs[x]

		// 切断
		tf := tails[f]
		tails[f] = n.prev
		n.prev.next = nil

		// 接続
		n.prev = tails[t]
		tails[t].next = n
		tails[t] = tf
	}
	containers := make([]int, N)

	for i := 0; i < len(heads); i++ {
		for p := heads[i].next; p != nil; p = p.next {
			containers[p.no] = i + 1
		}
	}
	for i := 0; i < N; i++ {
		fmt.Println(containers[i])
	}
}

type node struct {
	no   int
	prev *node
	next *node
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
