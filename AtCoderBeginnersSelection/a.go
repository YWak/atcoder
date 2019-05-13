package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IntSlice(str string) []int {
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
	a := IntSlice(line1)[0]

	// 二行目
	stdin.Scan()
	line2 := stdin.Text()
	ints2 := IntSlice(line2)
	b := ints2[0]
	c := ints2[1]

	// 三行目
	stdin.Scan()
	line3 := stdin.Text()

	fmt.Printf("%d %s\n", a+b+c, line3)
}
