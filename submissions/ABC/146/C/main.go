package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A := nextInt64()
	B := nextInt64()
	X := nextInt64()

	left := int64(-1)       // 買える
	right := int64(1e9 + 1) // 買えない

	for left+1 < right {
		mid := (left + right) / 2

		// mid is ok
		ok := A*mid+B*digits(mid) <= X
		if ok {
			left = mid
		} else {
			right = mid
		}
	}

	if left < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(left)
	}
}

func digits(n int64) int64 {
	for i := 1; true; i++ {
		if n < 10 {
			return int64(i)
		}
		n /= 10
	}
	return -1
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
