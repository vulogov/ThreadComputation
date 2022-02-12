package ThreadComputation

import (
	"testing"
)

func TestVars1(t *testing.T) {
	tc := Init()
	tc = tc.Eval("42 $answer")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestVars2(t *testing.T) {
	tc := Init()
	tc = tc.Eval("42 $answer answer")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("variable retrieval is failed: %v", res)
	}
}
