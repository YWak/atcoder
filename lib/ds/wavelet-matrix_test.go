package ds_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ywak/atcoder/lib/ds"
)

func TestWaveletMatrix(t *testing.T) {
	size := 100
	for q := 0; q < 100; q++ {
		arr := make([]int, size)
		for i := range arr {
			arr[i] = rand.Intn(1 << 20)
		}
		t.Log(arr)
		wm := ds.NewWaveletMatrix(arr)

		t.Run(fmt.Sprintf("Get(%d)", q), func(t *testing.T) {
			for i, expected := range arr {
				actual := wm.Get(i)
				if actual != expected {
					t.Errorf("Get(%d) = %d, want %d", i, actual, expected)
				}
			}
		})
	}
}
