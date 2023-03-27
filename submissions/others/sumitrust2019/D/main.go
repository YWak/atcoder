package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	S := nextBytes()

	hist1 := map[int]bool{}
	hist2 := map[int]bool{}
	hist3 := map[int]bool{}

	for i := 0; i < N; i++ {
		n := int(S[i] - '0')

		for h2 := range hist2 {
			hist3[h2*10+n] = true
		}

		for h1 := range hist1 {
			hist2[h1*10+n] = true
		}

		hist1[n] = true
	}

	fmt.Println(len(hist3))
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
