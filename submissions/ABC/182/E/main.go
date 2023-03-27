package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var LIGHT int = 2
var BLOCK int = -1
var BLIGHT int = 1
var DARK int = 0

func solve(area [][]int, H, W int) int {
	var on bool
	for h := 0; h < H; h++ {

		// 左
		on = false
		for w := 0; w < W; w++ {
			if area[h][w] == BLOCK {
				on = false
			} else if area[h][w] == LIGHT {
				on = true
			} else if on {
				area[h][w] = BLIGHT
			}
		}

		// 右
		on = false
		for w := W - 1; w >= 0; w-- {
			if area[h][w] == BLOCK {
				on = false
			} else if area[h][w] == LIGHT {
				on = true
			} else if on {
				area[h][w] = BLIGHT
			}
		}
	}
	for w := 0; w < W; w++ {
		// 上
		on = false
		for h := 0; h < H; h++ {
			if area[h][w] == BLOCK {
				on = false
			} else if area[h][w] == LIGHT {
				on = true
			} else if on {
				area[h][w] = BLIGHT
			}
		}
		on = false
		for h := H - 1; h >= 0; h-- {
			if area[h][w] == BLOCK {
				on = false
			} else if area[h][w] == LIGHT {
				on = true
			} else if on {
				area[h][w] = BLIGHT
			}
		}
	}
	c := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if area[h][w] > 0 {
				c++
			}
		}
	}
	return c
}

func main() {
	H := nextInt()
	W := nextInt()
	N := nextInt()
	M := nextInt()

	area := make([][]int, H)
	for i := 0; i < H; i++ {
		area[i] = make([]int, W)
	}

	for i := 0; i < N; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		area[a][b] = LIGHT
	}
	for i := 0; i < M; i++ {
		c := nextInt() - 1
		d := nextInt() - 1
		area[c][d] = BLOCK
	}

	fmt.Println(solve(area, H, W))
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
