package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A := nextInt()
	B := nextInt()

	for i := 1; true; i++ {
		if int(float64(i)*0.08) == A && int(float64(i)*0.1) == B {
			fmt.Println(i)
			return
		}

		if int(float64(i)*0.08) > A || int(float64(i)*0.1) > B {
			break
		}
	}

	fmt.Println(-1)
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
