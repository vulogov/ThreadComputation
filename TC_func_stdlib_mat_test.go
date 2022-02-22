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

func TestMat3(t *testing.T) {
	tc := Init()
	tc = tc.Eval("M[2 2] matrix.Set[1 1 3.14] print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestMat4(t *testing.T) {
	tc := Init()
	tc = tc.Eval("M[2 2] matrix.Set[1 1 3.14] matrix.Get[1 1]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3.14" {
		t.Fatalf("M[2 2] matrix.Set[1 1 3.14] matrix.Get[1 1] failed: %v", res)
	}
}

func TestMat5(t *testing.T) {
	tc := Init()
	tc = tc.Eval("M[2 2] matrix.Make[1 2 3 4] print matrix.Get[1 1]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4" {
		t.Fatalf("M[2 2] matrix.Make[1 2 3 4] print matrix.Get[1 1] failed: %v", res)
	}
}

func TestMat6(t *testing.T) {
	tc := Init()
	tc = tc.Eval("6 5 4 3 2 1 M[2 3] matrix.Make print matrix.Get[1 0]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4" {
		t.Fatalf("6 5 4 3 2 1 M[2 3] matrix.Make print matrix.Get[1 0] failed: %v", res)
	}
}
