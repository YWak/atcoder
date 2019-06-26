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
	C := nextInt64()
	D := nextInt64()

	if D < C {
		D, C = C, D
	}

	fmt.Println(count(B, C, D) - count(A-1, C, D))
}

func count(n, a, b int64) int64 {
	if b%a == 0 {
		return n - n/a
	}
	return n - n/a - n/b + n/lcm(a, b)
}

func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

func gcd(a, b int64) int64 {
	if b > a {
		return gcd(b, a)
	}

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
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
