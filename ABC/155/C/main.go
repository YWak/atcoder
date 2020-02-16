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
	hist := map[string]int{}

	for i := 0; i < N; i++ {
		S := nextString()
		n, ok := hist[S]

		if ok {
			n++
		} else {
			n = 1
		}
		hist[S] = n
	}
	c := 0
	ans := make(strings, 0)

	for s, n := range hist {
		if c > n {
			continue
		}
		if c < n {
			ans = make(strings, 0)
			c = n
		}
		ans = append(ans, s)
	}

	sort.Sort(ans)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}

type strings []string

func (a strings) Len() int           { return len(a) }
func (a strings) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a strings) Less(i, j int) bool { return a[i] < a[j] }
func max(a, b int) int {
	if a > b {
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
