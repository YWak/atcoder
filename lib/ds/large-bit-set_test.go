package ds_test

import (
	"testing"

	"github.com/ywak/atcoder/lib/ds"
)

func TestLargeBitSet(t *testing.T) {
	set := ds.NewLargeBitSet()

	for i := 0; i < 128; i++ {
		set.Add(i)
	}

	if !set.Has(1) {
		t.Errorf("set.Has(1) is not resolved")
	}

	if got := set.Before(0); got != nil {
		t.Errorf("set.Before(1) is not resolved. %v", got)
	}
	if got := set.Before(128); got == nil || *got != 127 {
		t.Errorf("set.Before(128) is not resolved. %v", got)
	}

	if got := set.After(1); got == nil || *got != 2 {
		t.Errorf("set.After(1) is not resolved. %v", got)
	}

	set.Add(1 << 30)
	set.Add(1 << 58)
}
