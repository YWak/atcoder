package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S := nextBytes()
	k := len(S)

	sum := 0
	countmod3 := [3]int{}

	for i := 0; i < k; i++ {
		n := int(S[i] - '0')
		sum = (sum + n) % 3
		countmod3[n%3]++
	}

	sm := sum % 3
	c := 0
	if sm == 0 {
		c = 0
	} else {
		if countmod3[sm] >= 1 {
			c = 1
		} else if countmod3[3-sm] >= 2 {
			c = 2
		} else {
			c = -1
		}
		if c >= k {
			c = -1
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

func nextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return a
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
