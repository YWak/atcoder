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
	length := ints1[0]
	n := ints1[1]

	stdin.Scan()
	line2 := stdin.Text()

	var head string
	var tail string

	if n == 1 {
		head = ""
	} else {
		head = line2[:n-1]
	}
	c := strings.ToLower(line2[n-1 : n])
	if n == length {
		tail = ""
	} else {
		tail = line2[n:]
	}

	fmt.Printf("%s%s%s\n", head, c, tail)
}
