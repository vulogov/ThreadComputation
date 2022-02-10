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
