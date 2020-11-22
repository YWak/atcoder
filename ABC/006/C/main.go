package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	M := nextInt()

	for i := 0; i <= N; i++ {
		for j := 0; j < 2; j++ {
			k := N - i - j

			if k < 0 {
				break
			}

			if i*2+j*3+k*4 == M {
				fmt.Println(i, j, k)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
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

// 遅いから極力使わない。
func nextBytes() []byte {
	return []byte(nextString())
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

func toi(b byte) int {
	return int(b - '0')
}

func nextLongIntAsArray() []int {
	s := nextString()
	l := len(s)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = toi(s[i])
	}
	return arr
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
