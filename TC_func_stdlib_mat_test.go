package ThreadComputation

import (
	"testing"
)

func TestMat1(t *testing.T) {
	tc := Init()
	tc = tc.Eval("m print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestMat2(t *testing.T) {
	tc := Init()
	tc = tc.Eval("M[2 2] print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
