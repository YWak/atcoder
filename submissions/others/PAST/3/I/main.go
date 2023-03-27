package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	inv := false
	row := make([]int64, N)
	col := make([]int64, N)
	for i := 0; i < N; i++ {
		row[i] = int64(i)
		col[i] = int64(i)
	}

	Q := nextInt()
	for i := 0; i < Q; i++ {
		q := nextInt()

		if q == 1 {
			a := nextInt() - 1
			b := nextInt() - 1
			if inv {
				row[a], row[b] = row[b], row[a]
			} else {
				col[a], col[b] = col[b], col[a]
			}
		} else if q == 2 {
			a := nextInt() - 1
			b := nextInt() - 1
			if inv {
				col[a], col[b] = col[b], col[a]
			} else {
				row[a], row[b] = row[b], row[a]
			}
		} else if q == 3 {
			inv = !inv
		} else {
			a := nextInt() - 1
			b := nextInt() - 1

			if inv {
				fmt.Println(int64(N)*col[b] + row[a])
			} else {
				fmt.Println(int64(N)*col[a] + row[b])
			}
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
