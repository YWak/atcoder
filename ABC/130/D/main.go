package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func print(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%2d ", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	N := nextInt()
	K := nextInt64()
	sum := make([]int64, N+1) // sum[i+1]はa[0]からa[i]までの合計
	c := int64(0)
	p := 0

	for i := 0; i < N; i++ {
		a := nextInt64()
		si := i + 1
		sum[si] = sum[i] + a

		// sum[si] - sum[p] >= K なら、p以下の要素を足しても題意を満たす
		for p <= si {
			if sum[si]-sum[p] < K {
				c += int64(p)
				break
			}
			p++
		}
	}

	fmt.Println(c)
}
