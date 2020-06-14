package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	A := make([]int, N)
	numbers := map[int][]int{}
	M := 1

	for i := 0; i < N; i++ {
		a := nextInt()
		A[i] = a
		numbers[a] = append(numbers[a], i)
		M = max(a, M)
	}

	c := 0

	for i, l := range numbers {
		if len(l) > 1 {
			continue
		}
		divs := divisors(i)
		// debug(i, divs)
		exists := false
		for j := 0; j < len(divs); j++ {
			_, ok := numbers[divs[j]]
			if ok && divs[j] != i {
				exists = true
				break
			}
		}
		if !exists {
			c++
		}
	}

	fmt.Println(c)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func divisors(n int) []int {
	divs := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)

			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}

	return divs
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

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args)
}
