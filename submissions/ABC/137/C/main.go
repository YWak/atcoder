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
	d := map[string]int{}

	for i := 0; i < N; i++ {
		var t text = nextBytes()
		sort.Sort(t)
		s := string(t)

		before, ok := d[s]

		if !ok {
			d[s] = 1
		} else {
			d[s] = before + 1
		}
	}

	count := int64(0)
	for _, c := range d {
		count += int64(c) * int64(c-1) / 2
	}

	fmt.Println(count)
}

type text []byte

func (a text) Len() int           { return len(a) }
func (a text) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a text) Less(i, j int) bool { return a[i] < a[j] }

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
