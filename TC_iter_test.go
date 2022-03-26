package ThreadComputation

import (
	"testing"
)

func TestIter1(t *testing.T) {
	tc := Init()
	i := tc.Iterator(int64(42))
	n := i.Next()
	if n.(int64) != 0 {
		t.Errorf("Int iterator had failed: %v", n)
	}
}

func TestIter2(t *testing.T) {
	tc := Init()
	i := tc.Iterator(int64(42))
	n := i.Prev()
	if n.(int64) != 42 {
		t.Errorf("Int iterator had failed: %v", n)
	}
}
