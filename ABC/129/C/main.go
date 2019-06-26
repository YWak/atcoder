package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = initStdin()

func initStdin() *bufio.Scanner {
	var stdin = bufio.NewScanner(os.Stdin)
	stdin.Split(bufio.ScanWords)
	return stdin
}

func nextString() string {
	stdin.Scan()
	return stdin.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func nextInt64() int64 {
	i, _ := strconv.ParseInt(nextString(), 10, 64)

	return i
}

func main() {
	n := nextInt()
	m := nextInt()

	broken := make([]int, n+1)
	patterns := make([]int, n+1)

	for i := 0; i < m; i++ {
		pos := nextInt()
		broken[pos] = 1
	}

	patterns[0] = 1

	if broken[1] == 1 {
		patterns[1] = 0
	} else {
		patterns[1] = 1
	}

	for i := 2; i <= n; i++ {
		if broken[i] == 1 {
			// 壊れていれば通れない
			patterns[i] = 0
		} else {
			patterns[i] = (patterns[i-1] + patterns[i-2]) % 1000000007
		}
	}

	fmt.Println(patterns[n])
}
