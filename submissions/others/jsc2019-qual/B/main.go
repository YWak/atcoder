package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	K := nextInt64()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	c1 := int64(0)
	c2 := int64(0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if A[i] > A[j] {
				if i < j {
					c1++
				}
				c2++
			}
		}
	}

	com := mdiv(mmul(K, K-1), 2)
	fmt.Println(madd(mmul(c1, int64(K)), mmul(c2, com)))
}

var mod = int64(1e9 + 7)

// mod を法とする加算
func madd(a, b int64) int64 { return (a + b) % mod }

// mod を法とする減算
func msub(a, b int64) int64 { return (a - b + mod) % mod }

// mod を法とする乗算
func mmul(a, b int64) int64 { return (a * b) % mod }

// mod を法とする除算
func mdiv(a, b int64) int64 { return mmul(a, minv(b)) }

// mod を法とした逆元
func minv(a int64) int64 {
	// 拡張ユークリッドの互除法
	b := mod
	u := int64(1)
	v := int64(0)
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod

	if u < 0 {
		u += mod
	}
	return u
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
