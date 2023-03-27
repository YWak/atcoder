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
	V := make(values, N)

	for i := 0; i < N; i++ {
		V[i] = float64(nextInt())
	}

	// vをソートして前から使っていく
	sort.Sort(V)
	ans := V[0]

	for n := 1; n < N; n++ {
		ans = (ans + V[n]) / 2
	}

	fmt.Println(ans)
}

type values []float64

func (a values) Len() int           { return len(a) }
func (a values) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a values) Less(i, j int) bool { return a[i] < a[j] }

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
