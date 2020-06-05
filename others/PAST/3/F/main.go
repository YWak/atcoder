package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()

	S := make([]byte, N)
	mat := make([][]byte, N)
	candidates := make([]map[byte]int, N)

	for i := 0; i < N; i++ {
		bytes := nextBytes()
		mat[i] = bytes
		candidates[i] = map[byte]int{}

		for j := 0; j < N; j++ {
			candidates[i][bytes[j]] = 1
		}
	}

	for i := 0; i < N/2; i++ {
		r := N - 1 - i
		found := false
		for j := 0; j < N; j++ {
			c := mat[i][j]
			_, ok := candidates[r][c]

			if ok {
				S[i] = c
				S[r] = c
				found = true
				break
			}
		}
		if !found {
			fmt.Println(-1)
			return
		}
	}
	if N%2 == 1 {
		mid := (N / 2)
		S[mid] = mat[mid][0]
	}
	fmt.Println(string(S))
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
