package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := nextBytes()
	L := len(N)
	c := 0
	k := 0

	for i := L - 1; i >= 0; i-- {
		n := int(N[i] - '0')

		if n <= 5 {
			// 0,1,2,3,4,5ならそのまま払う
			c += n
			k = 0
		} else {
			// 6,7,8,9なら1枚払って(10-n)枚もらう
			c += 1 + (10 - n)

			// 直前も繰り上げなら(連続している分)枚減る
			if k > 0 {
				c -= k + 1
			}
			k++
		}
	}
	if k > 0 {
		c -= k - 1
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
