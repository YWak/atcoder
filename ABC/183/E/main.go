package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	H := nextInt()
	W := nextInt()

	S := make([][]byte, H)
	for i := 0; i < H; i++ {
		S[i] = nextBytes()
	}

	dp := [2000][2000]mint{}
	dp[0][0] = 1

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if S[h][w] == '#' {
				continue
			}

			// 縦
			for i := h + 1; i < H; i++ {
				if S[i][w] == '#' {
					break
				}
				dp[i][w] = dp[i][w].add(dp[h][w])
			}
			// 横
			for i := w + 1; i < W; i++ {
				if S[h][i] == '#' {
					break
				}
				dp[h][i] = dp[h][i].add(dp[h][w])
			}
			// 斜め
			for i := 1; h+i < H && w+i < W; i++ {
				if S[h+i][w+i] == '#' {
					break
				}
				dp[h+i][w+i] = dp[h+i][w+i].add(dp[h][w])
			}
		}
	}

	fmt.Println(dp[H-1][W-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type mint int

const mod = mint(1e9 + 7)

// add は a + bを返します
func (a mint) add(b mint) mint {
	return (a + b) % mod
}

// sub は a - bを返します
func (a mint) sub(b mint) mint {
	return (a - b + mod) % mod
}

// mul は a * bを返します
func (a mint) mul(b mint) mint {
	return (a * (b % mod)) % mod
}

// div は a/bを返します
func (a mint) div(b mint) mint {
	return a.mul(b.inv())
}

// inv は aの逆元を返します
func (a mint) inv() mint {
	// 拡張ユークリッドの互除法
	b := mod
	u := mint(1)
	v := mint(0)
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	return (u + mod) % mod
}

// pow は a ^ bを返します
func (a mint) pow(b mint) mint {
	ans := mint(1)

	for b > 0 {
		if b&1 == 1 {
			ans = ans.mul(a)
		}
		a = a.mul(a)
		b = b >> 1
	}
	return ans
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
