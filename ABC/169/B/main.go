package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	m := uint64(1)
	N := nextInt()

	A := make([]uint64, N)

	max := uint64(1e18)

	for i := 0; i < N; i++ {
		A[i] = uint64(nextInt64())
		if A[i] == 0 {
			fmt.Println(0)
			return
		}
	}

	for i := 0; i < N; i++ {
		ma := m * A[i]
		if ma < m || ma > max {
			fmt.Println(-1)
			return
		}
		m = ma
	}

	fmt.Println(m)
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
