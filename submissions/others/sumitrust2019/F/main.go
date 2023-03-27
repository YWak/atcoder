package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	T1 := nextInt64()
	T2 := nextInt64()
	A1 := nextInt64()
	A2 := nextInt64()
	B1 := nextInt64()
	B2 := nextInt64()

	P := (A1 - B1) * T1
	Q := (A2 - B2) * T2

	if P > 0 {
		P = -P
		Q = -Q
	}

	if P+Q <= 0 {
		if P+Q == 0 {
			fmt.Println("infinity")
		} else {
			fmt.Println(0)
		}
		return
	}

	// 追いつく回数を求める
	s := (-P) / (P + Q)
	t := (-P) % (P + Q)

	if t == 0 {
		fmt.Println(s * 2)
	} else {
		fmt.Println(s*2 + 1)
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
