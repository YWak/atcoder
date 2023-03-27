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

	a := make([]int, M)
	for i := 0; i < M; i++ {
		a[i] = nextInt()
	}

	sushi := make([]int, M)
	for i := 0; i < M; i++ {
		sushi[i] = -1
	}

	children := make([]int, N)
	for i := 0; i < M; i++ {
		ng := -1
		ok := N

		for abs(ok-ng) > 1 {
			mid := (ok + ng) / 2

			if children[mid] < a[i] {
				ok = mid
			} else {
				ng = mid
			}
		}

		if ok == N {
			// 誰も食べない
			continue
		}
		children[ok] = a[i]
		sushi[i] = ok + 1
	}

	for i := 0; i < M; i++ {
		fmt.Println(sushi[i])
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
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
