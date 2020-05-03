package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n = 100 // i^5 > math.MaxInt32 となるi

func main() {
	X := nextInt64()
	f := make([]int64, n)

	for i := 0; i < n; i++ {
		f[i] = int64(i*i) * int64(i*i*i)
	}
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			// (+i)^5 - (+j)^5
			if f[i]-f[j] == X {
				fmt.Printf("%d %d\n", i, j)
				return
			}
			// (+i)^5 - (-j)^5
			if f[i]+f[j] == X {
				fmt.Printf("%d %d\n", i, -j)
				return
			}
		}
	}
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
