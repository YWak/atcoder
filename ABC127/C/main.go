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
	n := ints1[0]
	m := ints1[1]

	lMax := 0
	rMin := n

	// ゲートはm個
	for i := 0; i < m; i++ {
		stdin.Scan()
		line := stdin.Text()
		ints := intSlice(line)

		if lMax < ints[0] {
			lMax = ints[0]
		}
		if rMin > ints[1] {
			rMin = ints[1]
		}
	}

	if lMax > rMin {
		fmt.Println(0)
	} else {
		fmt.Println(rMin - lMax + 1)
	}
}
