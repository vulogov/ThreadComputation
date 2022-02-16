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

func TestVars11(t *testing.T) {
	tc := Init()
	tc = tc.Eval("tc.Version")
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

func TestVars3(t *testing.T) {
	tc := Init()
	tc.SetVariable("answer", 42)
  tc = tc.Eval("answer")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("set local variable not working: %v", res)
	}
}

func TestVars4(t *testing.T) {
	tc := Init()
	SetVariable("answer", 42)
  tc = tc.Eval("answer")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("set global variable not working: %v", res)
	}
}

func TestVars6(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 local[\"answer\"] answer")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("set local variable not working: %v", res)
	}
}

func TestVars7(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 global[\"answer\"]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	tc1 := Init()
	tc1 = tc1.Eval("answer")
	if tc1.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc1.GetAsString()
	if res != "42" {
		t.Fatalf("set global variable not working: %v", res)
	}
}
