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

	n := gcd(A, B)
	c := 1

	if n%2 == 0 {
		c++

		for n%2 == 0 {
			n >>= 1
		}
	}
	for i := int64(3); n > 1; i += 2 {
		if n%i != 0 {
			continue
		}
		c++

		for n%i == 0 {
			n /= i
		}
	}

	fmt.Println(c)
}

func gcd(a, b int64) int64 {
	if b > a {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func root(a int64) int64 {
	i := int64(1)
	for i*i <= a {
		i++
	}
	return i + 1
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
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
