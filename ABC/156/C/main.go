package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := nextInt()
	X := make(points, N)

	for i := 0; i < N; i++ {
		X[i] = nextInt()
	}
	sort.Sort(X)

	dmin := math.MaxInt32

	for i := 0; i <= 100; i++ {
		d := 0
		for j := 0; j < N; j++ {
			xp := (X[j] - i)
			d += xp * xp
		}
		if dmin > d {
			dmin = d
		}
	}

	fmt.Println(dmin)
}

type points []int

func (a points) Len() int           { return len(a) }
func (a points) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a points) Less(i, j int) bool { return a[i] < a[j] }

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
