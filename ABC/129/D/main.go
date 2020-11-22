package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	return []byte(nextString())
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func nextInt64() int64 {
	i, _ := strconv.ParseInt(nextString(), 10, 64)

	return i
}

func print(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%2d ", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	H := nextInt()
	W := nextInt()

	objects := make([][]byte, H)
	L := make([][]int, H)
	R := make([][]int, H)
	D := make([][]int, H)
	U := make([][]int, H)

	for i := 0; i < H; i++ {
		objects[i] = nextBytes()
		L[i] = make([]int, W)
		R[i] = make([]int, W)
		D[i] = make([]int, W)
		U[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			k := H - 1 - i
			l := W - 1 - j

			if objects[i][j] == '.' {
				if i == 0 {
					U[i][j] = 1
				} else {
					U[i][j] = U[i-1][j] + 1
				}
				if j == 0 {
					L[i][j] = 1
				} else {
					L[i][j] = L[i][j-1] + 1
				}
			}

			if objects[k][l] == '.' {
				if k == H-1 {
					D[k][l] = 1
				} else {
					D[k][l] = D[k+1][l] + 1
				}
				if l == W-1 {
					R[k][l] = 1
				} else {
					R[k][l] = R[k][l+1] + 1
				}
			}
		}
	}

	max := 0

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			m := U[i][j] + R[i][j] + D[i][j] + L[i][j] - 3
			if max < m {
				max = m
			}
		}
	}

	//print(u)
	//print(r)
	//print(d)
	//print(l)
	fmt.Println(max)
}
