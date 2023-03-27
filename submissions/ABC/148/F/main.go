package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var T [][]int
var dist []int

func main() {
	N := nextInt()
	u := nextInt() - 1
	v := nextInt() - 1
	T = make([][]int, N)
	dist = make([]int, N)

	for i := 0; i < N-1; i++ {
		A := nextInt() - 1
		B := nextInt() - 1
		T[A] = append(T[A], B)
		T[B] = append(T[B], A)
	}

	// 高橋くんを根として距離を出す
	dist[u] = 1
	fillDist(u)

	// 青木くんは最短を目指すので、とりあえず近づく
	var vk int
	var k int
	if (dist[v]-dist[u])%2 == 1 {
		k = (dist[v] - dist[u]) / 2
		vk = findUpper(v, k)
	} else {
		k = (dist[v] - dist[u] - 1) / 2
		vk = findUpper(v, k)
	}
	uk := findUpper(v, k+1)

	// ukから見て一番遠いところまでの距離。ただし、vkは通れない
	l := dfs(uk, vk)

	fmt.Println(k + l - 1)
}

func dfs(r, ignore int) int {
	m := 0
	for i := 0; i < len(T[r]); i++ {
		if T[r][i] == ignore {
			continue
		}
		m = max(m, dfs(T[r][i], r))
	}
	return m + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findUpper(r, d int) int {
	if d == 0 {
		return r
	}

	for i := 0; i < len(T[r]); i++ {
		if dist[T[r][i]] < dist[r] {
			return findUpper(T[r][i], d-1)
		}
	}
	return -1
}

func fillDist(r int) {
	for i := 0; i < len(T[r]); i++ {
		n := T[r][i]
		if dist[n] == 0 {
			dist[T[r][i]] = dist[r] + 1
			fillDist(n)
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
