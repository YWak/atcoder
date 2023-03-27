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
	K := nextInt64()
	A := make([]int64, N)
	sum := int64(0)

	for i := 0; i < N; i++ {
		A[i] = nextInt64()
		sum += A[i]
	}

	divs := dividersOf(sum)
	l := len(divs)
	for i := 0; i < l; i++ {
		div := divs[l-1-i]

		mods := make(ints, N)
		for i := 0; i < N; i++ {
			mods[i] = A[i] % div
		}
		sort.Sort(mods)

		d := int64(0)
		k := int64(0)
		l := 0
		r := N - 1

		// fmt.Println(mods)
		for l <= r {
			if d < 0 {
				// fmt.Println(div, "r", r, div-mods[r])
				d += div - mods[r]
				r--
			} else {
				// fmt.Println(div, "l", l, -mods[l])
				d -= mods[l]
				k += mods[l]
				l++
			}
		}

		if k <= K && d == 0 {
			fmt.Println(div)
			return
		}
	}
}

type ints []int64

func (a ints) Len() int           { return len(a) }
func (a ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ints) Less(i, j int) bool { return a[i] < a[j] }

func dividersOf(n int64) []int64 {
	div1 := make([]int64, 0)
	div2 := make([]int64, 0)

	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			div1 = append(div1, i)

			if i != n/i {
				div2 = append(div2, n/i)
			}
		}
	}
	l := int64(len(div2))
	for i := int64(0); i < l; i++ {
		div1 = append(div1, div2[l-1-i])
	}
	return div1
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
