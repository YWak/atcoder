package main

import (
	"bufio"
	"fmt"
	"math"
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
	k int
	s []int
	p int
}

func refbit(i int, b uint) int {
	return (i >> b) & 1
}

func main() {
	N := nextInt()
	M := nextInt()

	lights := []pair{}

	// 初期化
	for i := 0; i < M; i++ {
		k := nextInt()
		lights = append(lights, pair{k, []int{}, 0})

		for j := 0; j < k; j++ {
			lights[i].s = append(lights[i].s, nextInt())
		}
	}
	for i := 0; i < M; i++ {
		lights[i].p = nextInt()
	}

	var c int
	// fmt.Println(lights)

	states := int(math.Pow(float64(2), float64(N)))
next:
	for n := 0; n < states; n++ {
		// fmt.Printf("%08b\n", n)
		for i := 0; i < M; i++ {
			ons := 0
			for j := 0; j < lights[i].k; j++ {
				ons += refbit(n, uint(lights[i].s[j]-1))
			}
			// fmt.Printf("%d[%d] ons = %d\n", n, i, ons)
			if ons%2 != lights[i].p {
				// fmt.Printf("%dで違反\n", i)
				continue next
			}
		}
		c++
	}

	fmt.Println(c)
}
