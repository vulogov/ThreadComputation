package ThreadComputation

import (
	"testing"
)

func TestInit(t *testing.T) {
	tc := Init()
	if tc.InAttr != false {
		t.Fatalf("ThreadComputation is failed to initialize")
	}
}

func TestEval1(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("42")
	if tc1.Errors() != 0 {
		t.Fatalf("ThreadComputation parse failed")
	}
}
