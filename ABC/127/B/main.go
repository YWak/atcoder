package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intSlice(str string) []int {
	splitted := strings.Split(str, " ")
	ret := make([]int, len(splitted))

	for i, str := range splitted {
		if strings.TrimSpace(str) != "" {
			n, _ := strconv.Atoi(str)
			ret[i] = n
		}
	}

	return ret
}

func main() {

	stdin := bufio.NewScanner(os.Stdin)

	// 一行目
	stdin.Scan()
	line1 := stdin.Text()
	ints1 := intSlice(line1)
	r := ints1[0]
	d := ints1[1]
	x := ints1[2]
	x0 := x

	for i := 0; i < 10; i++ {
		x1 := r*x0 - d
		fmt.Println(x1)
		x0 = x1
	}
}
