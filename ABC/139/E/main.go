package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var a [][]int

func main() {
	N := nextInt()
	M := N - 1
	a = make([][]int, N)
	nexts := make([]int, N)

	for i := 0; i < N; i++ {
		a[i] = make([]int, M)

		for j := 0; j < M; j++ {
			a[i][j] = nextInt() - 1
		}
	}

	// greedy
	var n int
	for n = 1; true; n++ {
		added := false
		match := make([]int, N)
		for i := 0; i < N; i++ {
			if nexts[i] == M {
				continue
			}

			j := a[i][nexts[i]] // iの次の対戦相手

			if i > j { // 対戦相手は順序付けられている
				continue
			}
			if nexts[j] == M {
				continue
			}
			if i != a[j][nexts[j]] { // 対戦相手の予定と一致していなければ無視
				continue
			}
			if match[i] == 1 || match[j] == 1 { // 試合済みなら無効
				continue
			}

			// fmt.Printf("%d: %dvs%d\n", n, i, j)
			nexts[i]++
			nexts[j]++
			match[i]++
			match[j]++
			added = true
		}

		if !added {
			break
		}
	}
	// 全員が試合したかどうかを判断する
	for i := 0; i < N; i++ {
		if nexts[i] != M {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(n - 1)
	return
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
