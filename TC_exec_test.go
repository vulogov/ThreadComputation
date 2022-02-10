package ThreadComputation

import (
	"testing"
)


func TestEval2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("\"Hello\"")
  if tc.Res.Len() == 0 {
		t.Fatalf("ThreadComputation parse failed")
	}
}

func TestEval21(t *testing.T) {
	tc := Init()
  tc = tc.Eval("\"Hello\"")
  res := tc.GetAsString()
  if res != "Hello" {
		t.Fatalf("ThreadComputation parse failed: %v ", res)
	}
}

func TestEval3(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42")
  if tc.Res.Len() == 0 {
		t.Fatalf("ThreadComputation parse failed")
	}
}

func TestEval31(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42")
  if tc.GetAsString() != "42" {
		t.Fatalf("ThreadComputation parse failed")
	}
}

func TestEval4(t *testing.T) {
	tc := Init()
  tc = tc.Eval("3.14")
  if tc.Res.Len() == 0 {
		t.Fatalf("ThreadComputation parse failed")
	}
}

func TestEval41(t *testing.T) {
	tc := Init()
  tc = tc.Eval("3.14")
  if tc.GetAsString() != "3.14" {
		t.Fatalf("ThreadComputation parse failed")
	}
}

func TestEval5(t *testing.T) {
	tc := Init()
  tc = tc.Eval("#TRUE")
  if tc.Res.Len() == 0 {
		t.Fatalf("ThreadComputation parse failed")
	}
}

func TestEval51(t *testing.T) {
	tc := Init()
  tc = tc.Eval("#TRUE")
  if tc.GetAsString() != "true" {
		t.Fatalf("ThreadComputation parse failed")
	}
}
