package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := nextInt()
	M := nextInt()

	H := make([]int, N)
	W := make([]int, M)

	for i := 0; i < N; i++ {
		H[i] = nextInt()
	}
	sort.Slice(H, func(i, j int) bool { return H[i] < H[j] })

	for i := 0; i < M; i++ {
		W[i] = nextInt()
	}
	sort.Slice(W, func(i, j int) bool { return W[i] < W[j] })

	ans := math.MaxInt64

	if N == 1 {
		for i := 0; i < M; i++ {
			ans = min(ans, abs(H[0]-W[i]))
		}
		fmt.Println(ans)
		return
	}

	// D1[n] は 前からn番目までのペアの累積和
	D1 := make([]int, N/2+1)
	// D2[n] は 後ろからn番目までのペアの累積和
	D2 := make([]int, N/2+1)

	D1[0] = 0
	D2[0] = 0

	for i := 1; i <= N/2; i++ {
		j := (i - 1) * 2
		D1[i] = D1[i-1] + abs(H[j]-H[j+1])
		D2[i] = D2[i-1] + abs(H[N-1-j]-H[N-2-j])
	}

	for w := 0; w < M; w++ {
		// 先生の位置 == W[w] が入って昇順になる位置。
		j := -1
		ng := N
		for abs(j-ng) > 1 {
			mid := (j + ng) / 2

			if H[mid] < W[w] {
				j = mid
			} else {
				ng = mid
			}
		}

		// 差分を取る
		c := 0
		// debug("j =", j)
		if (j+2)%2 == 1 {
			// 前半
			c += D1[(j+1)/2]
			c += abs(W[w] - H[j+1])
			c += D2[(N-j-1)/2]
		} else {
			c += D1[j/2]
			c += abs(W[w] - H[j])
			c += D2[(N-j)/2]
		}
		ans = min(ans, c)
	}

	fmt.Println(ans)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
