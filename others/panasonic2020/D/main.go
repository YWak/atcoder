package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	prev := make([]data, 0)
	curr := make([]data, 0)

	prev = append(prev, data{"a", 0})

	for i := 1; i < N; i++ {
		for j := 0; j < len(prev); j++ {
			d := prev[j]

			for k := 0; k <= d.max+1; k++ {
				curr = append(curr, data{d.str + chars[k], max(d.max, k)})
			}
		}
		prev = curr
		curr = make([]data, 0)
	}

	for i := 0; i < len(prev); i++ {
		fmt.Println(prev[i].str)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type data struct {
	str string
	max int
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
