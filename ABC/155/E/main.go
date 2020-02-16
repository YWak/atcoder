package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextBytes()

	c := 0
	d := 0
	k := 0

	for i := 0; i < len(N); i++ {
		n := int(N[i] - '0')

		// 0,1,2,3,4,5ならそのまま払う
		// 6,7,8,9なら1枚払って(10-n)枚もらう
		if n <= 5 {
			c += n
			d += k
			k = 0
		} else {
			c += 11 - n
			k++
		}
	}
	d += k - 1
	if d != 0 {
		d++
	}

	fmt.Println(c - d)
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
