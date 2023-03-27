package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	a := float64(nextInt())
	b := float64(nextInt())
	x := float64(nextInt())

	z := x / (a * a)
	deg := float64(180.0) / math.Pi

	if z >= b/2 {
		deg *= math.Atan2(b-z, a/2.0)
	} else {
		deg *= math.Atan2(b, (2.0*x)/(a*b))
	}
	fmt.Printf("%.10f\n", deg)
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
