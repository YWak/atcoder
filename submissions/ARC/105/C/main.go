package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := nextInt()
	M := nextInt()

	w := make([]int, N)
	maxw := 0
	minv := math.MaxInt64

	for i := 0; i < N; i++ {
		w[i] = nextInt()
		if maxw < w[i] {
			maxw = w[i]
		}
	}

	parts := make([]part, M)
	for i := 0; i < M; i++ {
		l := nextInt()
		v := nextInt()
		parts[i] = part{l, v}
		if minv > v {
			minv = v
		}
	}
	if maxw > minv {
		fmt.Println(-1)
		return
	}

	sort.Slice(parts, func(i, j int) bool {
		if parts[i].v != parts[j].v {
			return parts[i].v < parts[j].v
		}
		return parts[i].l < parts[j].l
	})

	perm := permutation(8, []int{}, 0)
	ans := math.MaxInt64

	for i := 0; i < len(perm); i++ {
		c := solve(parts, w, perm[i])
		if ans > c {
			ans = c
		}
	}

	fmt.Println(ans)
}

func solve(parts []part, w []int, order []int) int {
	// 全組み合わせの最小重量
	sumw := make([]int, len(w))
	sumw[0] = w[0]

	for i := 1; i < len(w); i++ {
		sumw[i] = sumw[i-1] + w[i]
	}

	return 0
}

func permutation(length int, curr []int, used int) [][]int {
	if len(curr) == length {
		return [][]int{curr}
	}

	ret := make([][]int, 0)

	for i := 0; i < length; i++ {
		if (used>>i)&1 == 1 {
			continue
		}
		next := append([]int{}, curr...)
		next = append(next, i)

		ret = append(ret, permutation(length, next, used|1<<i)...)
	}

	return ret
}

type part struct {
	l int
	v int
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
	fmt.Fprintln(os.Stderr, args...)
}
