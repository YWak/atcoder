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
	W := nextInt()

	uses := make([]d, N*2)

	for i := 0; i < N; i++ {
		s := nextInt()
		t := nextInt()
		p := nextInt()

		uses[i] = d{s, p}
		uses[N+i] = d{t, -p}
	}
	sort.Slice(uses, func(i, j int) bool {
		a := uses[i]
		b := uses[j]
		return a.t < b.t || a.t == b.t && a.p < b.p
	})

	t := -1
	w := 0
	uses = append(uses, d{1000000, -1000000001})
	for i := 0; i < len(uses); i++ {
		if t != uses[i].t {
			if w > W {
				fmt.Println("No")
				return
			}
			t = uses[i].t
		}
		w += uses[i].p
	}

	fmt.Println("Yes")
}

type d struct {
	t int
	p int
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
