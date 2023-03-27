package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var size = 202
var W = 1000000

func main() {
	N := nextInt()
	X := nextInt()
	Y := nextInt()

	distances := make([]int, W+1)
	for i := 0; i < W; i++ {
		distances[i] = math.MaxInt32
	}
	for i := 0; i < N; i++ {
		x := nextInt()
		y := nextInt()
		distances[conv(x, y)] = -1
	}
	dir := []node{
		node{+1, +1, 0},
		node{+0, +1, 0},
		node{-1, +1, 0},
		node{+1, +0, 0},
		node{-1, +0, 0},
		node{+0, -1, 0},
	}
	queue := make([]node, 0)
	queue = append(queue, node{0, 0, 0})
	distances[conv(0, 0)] = 0

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n.x == X && n.y == Y {
			fmt.Println(n.score)
			return
		}

		for i := 0; i < len(dir); i++ {
			d := dir[i]
			x, y := n.x+d.x, n.y+d.y
			if x < -size || x > size || y < -size || y > size {
				continue
			}

			w := conv(x, y)
			if distances[w] == -1 {
				continue
			}
			if distances[w] <= n.score+1 {
				continue
			}
			queue = append(queue, node{x, y, n.score + 1})
			distances[w] = n.score + 1
		}
	}

	fmt.Println(-1)
}

type node struct {
	x     int
	y     int
	score int
}

func conv(x, y int) int {
	x1 := x + 500
	y1 := y + 500

	return x1*1000 + y1
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
