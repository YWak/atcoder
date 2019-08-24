package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	M := nextInt()
	D := nextInt()

	c := 0

	for m := 1; m <= M; m++ {
		for d := 20; d <= D; d++ {
			d10 := d / 10
			d1 := d % 10
			if d1 >= 2 && d10 >= 2 && d10*d1 == m {
				c++
			}
		}
	}

	fmt.Println(c)
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
