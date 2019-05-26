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

type pair struct {
	name  string
	point int
	rank1 int
	rank2 int
}

func comparePoint(a, b pair) bool {
	if a.name == b.name {
		return a.point > b.point
	}
	return a.name < b.name
}

func swap(shops *[]pair, i, j int) {
	s := (*shops)[i]
	(*shops)[i] = (*shops)[j]
	(*shops)[j] = s
}

func main() {
	n := nextInt()
	shops := []pair{}

	// 初期化
	for i := 0; i < n; i++ {
		s := pair{nextString(), nextInt(), i, -1}
		shops = append(shops, s)
	}

	// 順位付け
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !comparePoint(shops[i], shops[j]) {
				swap(&shops, i, j)
			}
		}
	}
	for i := 0; i < n; i++ {
		shops[i].rank2 = i + 1
	}
	// fmt.Println(shops)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if shops[i].rank1 > shops[j].rank1 {
				swap(&shops, i, j)
			}
		}
	}
	// fmt.Println(shops)

	for i := 0; i < n; i++ {
		fmt.Println(shops[i].rank2)
	}
}
