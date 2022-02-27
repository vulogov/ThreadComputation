package ThreadComputation

import (
	"testing"
)

func TestGR1(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("wait")
	if tc1.Errors() != 0 {
		t.Fatalf("wait parse failed")
	}
}

func TestGR2(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("spawn\\ 'Hello world from goroutine' print \\ wait")
	if tc1.Errors() != 0 {
		t.Fatalf("spawn\\ 'Hello world' print \\ wait parse failed")
	}
}
