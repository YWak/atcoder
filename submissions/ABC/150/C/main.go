package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	p := 0
	q := 0

	for i := 0; i < N; i++ {
		p = p*10 + nextInt()
	}
	for i := 0; i < N; i++ {
		q = q*10 + nextInt()
	}

	rest := make([]int, 0)
	for i := 0; i < N; i++ {
		rest = append(rest, i+1)
	}

	cand := perm(0, rest)
	pr := -1
	qr := -1
	for i := 0; i < len(cand); i++ {
		if p == cand[i] {
			pr = i
		}
		if q == cand[i] {
			qr = i
		}
	}

	fmt.Println(abs(pr - qr))
}

func perm(curr int, rest []int) []int {
	if len(rest) == 0 {
		return []int{curr}
	}

	next := make([]int, 0)
	queue := make([]int, 0)

	for i := 0; i < len(rest); i++ {
		rest2 := make([]int, 0)
		for j := 0; j < len(queue); j++ {
			rest2 = append(rest2, queue[j])
		}
		for j := i + 1; j < len(rest); j++ {
			rest2 = append(rest2, rest[j])
		}

		next = append(next, perm(curr*10+rest[i], rest2)...)
		queue = append(queue, rest[i])
	}

	return next
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		print(arr[i])
		print(", ")
	}
	println()
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
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
