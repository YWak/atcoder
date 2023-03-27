package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A := nextInt64()
	R := nextInt64()
	N := nextInt64()

	r, over := pow(R, N-1)
	if over {
		fmt.Println("large")
		return
	}

	r, over = mul(r, A)
	if over {
		fmt.Println("large")
	} else {
		fmt.Println(r)
	}
}

func pow(x, n int64) (int64, bool) {
	if n == 0 {
		return 1, false
	}

	if n%2 == 0 {
		x2, over := mul(x, x)

		if over {
			return -1, true
		}

		return pow(x2, n/2)
	}

	p, over := pow(x, n-1)
	if over {
		return -1, true
	}

	return mul(x, p)
}

func mul(a, b int64) (int64, bool) {
	limit := int64(1e9)

	if a > limit {
		return -1, true
	}
	if b > limit {
		return -1, true
	}
	if a*b > limit {
		return -1, true
	}

	return a * b, false
}

func nthbit(a int64, n int) int { return int((a >> uint(n)) & 1) }

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
