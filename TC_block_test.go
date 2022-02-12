package ThreadComputation

import (
	"testing"
)

func TestBlock1(t *testing.T) {
	tc := Init()
	tc = tc.Eval("( 1 2 3 )")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestBlock2(t *testing.T) {
	tc := Init()
	tc = tc.Eval("(1 2 3) 42 (3 4 5)")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("block not rotated properly: %v", res)
	}
}
