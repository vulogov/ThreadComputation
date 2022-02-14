package ThreadComputation

import (
	"testing"
)


func TestFunc1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("FunctionNotExists")
  if tc.Errors() == 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("stack[ stack[42]]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("stack[ stack[42]] not working: %v", res)
	}
}
