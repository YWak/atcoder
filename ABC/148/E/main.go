package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt64()

	if N%2 == 1 {
		fmt.Println(0)
		return
	}
	a5 := int64(0)
	n5 := int64(10)

	// 5の指数と対応
	// 10の倍数 + 50の倍数 + 250の倍数 + 1250の倍数 + ...
	for n5 <= N {
		a5 += N / n5
		n5 *= 5
	}

	fmt.Println(a5)
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
