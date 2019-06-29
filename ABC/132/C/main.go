package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type scores []int

func (s scores) Len() int {
	return len(s)
}
func (s scores) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	N := nextInt()
	D := make(scores, N)

	for i := 0; i < N; i++ {
		D[i] = nextInt()
	}
	sort.Sort(D)

	abc := D[N/2-1]
	arc := D[N/2]

	if abc == arc {
		fmt.Println(0)
	} else {
		fmt.Println(arc - abc)
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
