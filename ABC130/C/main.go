package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = initStdin()

func initStdin() *bufio.Scanner {
	var stdin = bufio.NewScanner(os.Stdin)
	stdin.Split(bufio.ScanWords)
	return stdin
}

func nextString() string {
	stdin.Scan()
	return stdin.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func nextInt64() int64 {
	i, _ := strconv.ParseInt(nextString(), 10, 64)

	return i
}

func main() {
	W := nextInt()
	H := nextInt()
	x := nextInt()
	y := nextInt()

	area := float64(W) * float64(H) / 2.0
	var multiple string

	if float64(x) == float64(W)/2 && float64(y) == float64(H)/2 {
		multiple = "1"
	} else {
		multiple = "0"
	}

	fmt.Printf("%.6f %s\n", area, multiple)
}
