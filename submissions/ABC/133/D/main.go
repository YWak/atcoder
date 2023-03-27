package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	A := make([]int64, N)

	x0 := int64(0)

	for i := 0; i < N; i++ {
		A[i] = nextInt64()

		if i%2 == 0 {
			x0 += A[i]
		} else {
			x0 -= A[i]
		}
	}

	fmt.Print(x0)
	x := x0

	for i := 0; i < N-1; i++ {
		xi := A[i]*2 - x
		x = xi
		fmt.Printf(" %d", xi)
	}

	fmt.Println()
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
