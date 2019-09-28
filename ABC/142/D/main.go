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

	N := gcd(A, B)
	divisors := map[int64]bool{}

	for i := int64(1); i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		if isPrime(i) {
			divisors[i] = true
		}
		if isPrime(N/i) && i < N/i {
			divisors[N/i] = true
		}
	}
	// fmt.Println(N, divisors)
	fmt.Println(len(divisors))
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
	return i
}

func isPrime(a int64) bool {
	if a == 2 {
		return true
	}
	if a%2 == 0 {
		return false
	}

	for i := int64(3); i*i <= a; i += 2 {
		if a%i == 0 {
			return false
		}
	}
	return true
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
