package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	P := make([]point, N)
	for i := 0; i < N; i++ {
		P[i] = point{nextInt(), nextInt()}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if isOnLine(P[i], P[j], P[k]) {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}

func isOnLine(a, b, c point) bool {
	if a.x == b.x {
		return b.x == c.x
	}
	if a.y == b.y {
		return b.y == c.y
	}

	return (a.y-b.y)*(a.x-c.x) == (a.x-b.x)*(a.y-c.y)
}

type point struct {
	x int
	y int
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

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
