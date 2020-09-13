package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := nextInt()

	points := make([]point, N)

	for i := 0; i < N; i++ {
		x := nextInt()
		y := nextInt()
		p := point{x - y, x + y}
		points[i] = p
	}

	d := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})
	d = max(d, abs(points[0].x-points[N-1].x))
	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y
	})
	d = max(d, abs(points[0].y-points[N-1].y))

	fmt.Println(d)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
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
