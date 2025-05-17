package ds_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ywak/atcoder/lib/ds"
)

func TestSuccinctIndexableDictionary(t *testing.T) {
	size := 100000
	for q := 0; q < 10; q++ {
		s := ds.NewSuccinctIndexableDictionary(size)
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(2)
			s.Set(i, arr[i])
		}
		t.Log(arr)

		s.Build()

		t.Run(fmt.Sprintf("Get(%d)", q), func(t *testing.T) {
			for i := 0; i < size; i++ {
				actual := s.Get(i)
				if actual != arr[i] {
					t.Errorf("Get(%d) = %d, want %d", i, actual, arr[i])
				}
			}
		})
		t.Run(fmt.Sprintf("Rank(%d)", q), func(t *testing.T) {
			cnt := [2]int{}

			for i := 0; i < size; i++ {
				r0 := s.Rank(i, 0)
				if r0 != cnt[0] {
					t.Errorf("Rank(0, %d) = %d, want %d", i, r0, cnt[0])
				}

				r1 := s.Rank(i, 1)
				if r1 != cnt[1] {
					t.Errorf("Rank(1, %d) = %d, want %d", i, r1, cnt[1])
				}

				cnt[arr[i]]++
			}
		})
		t.Run(fmt.Sprintf("Select(%d)", q), func(t *testing.T) {
			cnt := [2]int{}
			for i := 0; i < size; i++ {
				cnt[arr[i]]++ // iはi+1番目のarr[i]の位置

				s := s.Select(arr[i], cnt[arr[i]])
				if s != i+1 {
					t.Errorf("Select(%d, %d) = %d, want %d", arr[i], cnt[arr[i]], s, i+1)
				}
			}
		})
	}
}
