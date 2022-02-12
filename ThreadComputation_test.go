package ThreadComputation

import (
	"testing"
)

func TestInit(t *testing.T) {
	tc := Init()
	if tc.InAttr != 0 {
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

func TestEval7(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("1 42")
	if tc1.Errors() != 0 {
		t.Fatalf("ThreadComputation parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("data in stack is not in order: %v and shall be 42", res)
	}
}
