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
		a[i] = make([]int, N-1)

		for j := 0; j < N-1; j++ {
			a[i][j] = nextInt() - 1
			// fmt.Print(a[i][j])
			// fmt.Print(" ")
		}
		// fmt.Println()
	}

	// greedy
	var n int
	for n = 1; true; n++ {
		added := false
		for i := 0; i < N; i++ {
			if nexts[i] == M {
				continue
			}

			j := a[i][nexts[i]] // iの次の対戦相手

			if i > j || nexts[j] == M {
				continue
			}

			if i != a[j][nexts[j]] { // 対戦相手の予定と一致していなければ無視
				continue
			}
			// fmt.Printf("%dvs%d\n", i, j)
			nexts[i]++
			nexts[j]++
			added = true
		}

		if !added {
			for i := 0; i < N; i++ {
				if nexts[i] != M {
					fmt.Println(-1)
					return
				}
			}
			fmt.Println(n)
			return
		}
	}

	fmt.Println(n)
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
