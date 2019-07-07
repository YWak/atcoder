package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	D := nextInt()

	vector := make([][]int, N)

	for i := 0; i < N; i++ {
		vector[i] = make([]int, D)

		for j := 0; j < D; j++ {
			vector[i][j] = nextInt()
		}
	}

	c := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			d2 := float64(0)

			for k := 0; k < D; k++ {
				d2 += float64(vector[i][k]-vector[j][k]) * float64(vector[i][k]-vector[j][k])
			}

			d := int(math.Sqrt(d2))

			if d2 == float64(d*d) {
				c++
			}
		}
	}

	fmt.Println(c)
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
