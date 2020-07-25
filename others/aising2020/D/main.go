package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var l = 200001
var f = make([]int, l)

func main() {
	f[0] = 0
	f[1] = 1

	// n <= 200000 でf[n]を求める
	for i := 2; i < l; i++ {
		popcount := bits(uint(i))
		f[i] = f[i%popcount] + 1
	}

	N := nextInt()
	X := nextBytes()
	c := 0

	for i := 0; i < N; i++ {
		if X[i] == '1' {
			c++
		}
	}
	if c == 1 {
		for i := 0; i < N; i++ {
			if X[i] == '1' {
				fmt.Println(0) // Xi == 0
				continue
			}
			// X[i] == 0
			if X[N-1] == '1' {
				fmt.Println(f[1] + 1) // Xi mod 2 == 1
			} else {
				fmt.Println(1) // Xi mod 2 == 0
			}
		}
		return
	}

	rem02 := make([]int, N) // rem0[i] = 2^i % (c-1)
	rem12 := make([]int, N) // rem1[i] = 2^i % (c+1)
	rem0 := 0
	rem1 := 0
	c0 := c + 1
	c1 := c - 1
	for i := 0; i < N; i++ {
		x := int(X[i] - '0')
		rem0 = (rem0*2 + x) % c0
		rem1 = (rem1*2 + x) % c1

		if i == 0 {
			rem02[i] = 1
			rem12[i] = 1
		} else {
			rem02[i] = (rem02[i-1] * 2) % c0
			rem12[i] = (rem12[i-1] * 2) % c1
		}
	}

	for i := 0; i < N; i++ {
		j := N - 1 - i
		if X[i] == '0' {
			rem := (rem0 + rem02[j] + c0) % c0
			fmt.Println(f[rem] + 1)
		} else {
			rem := (rem1 - rem12[j] + c1) % c1
			fmt.Println(f[rem] + 1)
		}
	}
}
func bits(v uint) int {
	v = (v & 0x55555555) + (v >> 1 & 0x55555555)
	v = (v & 0x33333333) + (v >> 2 & 0x33333333)
	v = (v & 0x0f0f0f0f) + (v >> 4 & 0x0f0f0f0f)
	v = (v & 0x00ff00ff) + (v >> 8 & 0x00ff00ff)
	v = (v & 0x0000ffff) + (v >> 16 & 0x0000ffff)
	return int(v)
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
	fmt.Fprintln(os.Stderr, args)
}
