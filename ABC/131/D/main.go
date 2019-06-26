package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := nextInt()
	var tasks pairs = make([]pair, N)

	for i := 0; i < N; i++ {
		tasks[i] = pair{nextInt64(), nextInt64()}
	}

	sort.Sort(tasks)

	t := int64(0)

	for i := 0; i < N; i++ {
		t += tasks[i].left

		if t > tasks[i].right {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

type pair struct {
	left  int64
	right int64
}
type pairs []pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Less(i, j int) bool {
	return p[i].right < p[j].right || p[i].right == p[j].right && p[i].left < p[j].left
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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
