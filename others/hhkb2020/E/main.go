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

	for h := 0; h < H; h++ {
		S[h] = nextBytes()
	}

	solve(H, W, S)
}

func solve(H, W int, S [][]byte) {

	top := make([][]int, H)
	left := make([][]int, H)
	bottom := make([][]int, H)
	right := make([][]int, H)

	K := 0
	for h := 0; h < H; h++ {
		top[h] = make([]int, W)
		left[h] = make([]int, W)
		bottom[h] = make([]int, W)
		right[h] = make([]int, W)

		for w := 0; w < W; w++ {
			if S[h][w] == '.' {
				K++
			}
		}
	}

	for t := 0; t < H; t++ {
		for l := 0; l < W; l++ {
			// 左
			if S[t][l] == '#' {
				left[t][l] = 0
			} else if l > 0 {
				left[t][l] = left[t][l-1] + 1
			} else {
				left[t][l] = 1
			}
			// 上
			if S[t][l] == '#' {
				top[t][l] = 0
			} else if t > 0 {
				top[t][l] = top[t-1][l] + 1
			} else {
				top[t][l] = 1
			}
		}
	}
	for b := H - 1; b >= 0; b-- {
		for r := W - 1; r >= 0; r-- {
			// 右
			if S[b][r] == '#' {
				right[b][r] = 0
			} else if r < W-1 {
				right[b][r] = right[b][r+1] + 1
			} else {
				right[b][r] = 1
			}
			// 下
			if S[b][r] == '#' {
				bottom[b][r] = 0
			} else if b < H-1 {
				bottom[b][r] = bottom[b+1][r] + 1
			} else {
				bottom[b][r] = 1
			}
		}
	}

	c := mint(0)
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if S[h][w] == '#' {
				continue
			}
			// S[h][w]は4回数えられているので、3回引く
			// 光ってないケースは2^(K-s)パターン
			s := top[h][w] + left[h][w] + bottom[h][w] + right[h][w] - 3
			ss := mint(2).pow(s).sub(mint(1)).mul(mint(2).pow(K - s))
			// debug(c, ss, s, mint(2).pow(K-s))
			c = c.add(ss)
		}
	}
	fmt.Println(c)
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
func (a mint) pow(b int) mint {
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
	return []byte(nextString())
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
