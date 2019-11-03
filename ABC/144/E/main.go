package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var N int
var K int64
var A ints
var F ints

func main() {
	N = nextInt()
	K = nextInt64()

	A = make(ints, N)
	F = make(ints, N)

	for i := 0; i < N; i++ {
		A[i] = nextInt64()
	}
	sort.Sort(sort.Reverse(A))

	for i := 0; i < N; i++ {
		F[i] = nextInt64()
	}
	sort.Sort(F)

	fmt.Println(binarysearch())
}

func binarysearch() int64 {
	l := int64(-1)
	r := int64(math.MaxInt64)

	for l+1 < r {
		m := (r + l) / 2
		if ok(m) {
			r = m
		} else {
			l = m
		}
	}

	return r
}

func ok(n int64) bool {
	k := int64(0)

	for i := 0; i < N; i++ {
		k += max(0, A[i]-n/F[i])
	}

	return k <= K
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type ints []int64

func (a ints) Len() int           { return len(a) }
func (a ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ints) Less(i, j int) bool { return a[i] < a[j] }

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
