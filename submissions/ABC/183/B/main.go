package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	Sx := float64(nextInt())
	Sy := float64(nextInt())
	Gx := float64(nextInt())
	Gy := float64(nextInt())

	Gy = -Gy
	a := (Gy - Sy) / (Gx - Sx)
	ans := (a*Sx - Sy) / a
	fmt.Printf("%.10f\n", ans)
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

func nextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return a
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
