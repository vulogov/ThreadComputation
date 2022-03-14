package ThreadComputation

import (
	"testing"
)

func TestStdlibStack1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 len")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "1" {
    t.Fatalf("Not pushed to stack: %v", res)
  }
}

func TestStdlibStack2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("stack[1 2 42]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("Not pushed to stack: %v", res)
  }
}

func TestStdlibStack3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("stack[1 2 stack[3 4 42]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("Not pushed to stack: %v", res)
  }
}

func TestStdlibStack4(t *testing.T) {
  tc := Init()
  tc = tc.Eval("stack[1 2 3] clr len")
  res := tc.GetAsString()
  if  res != "0" {
    t.Fatalf("clr function not working: %v", res)
  }
}

func TestStdlibStack5(t *testing.T) {
  tc := Init()
  tc = tc.Eval("1 2 3 clr[42]")
  res := tc.GetAsString()
  if  res != "42" {
    t.Fatalf("clr function not working: %v", res)
  }
}

func TestStdlibStack6(t *testing.T) {
  tc := Init()
  tc = tc.Eval("42 dup len")
  res := tc.GetAsString()
  if  res != "2" {
    t.Fatalf("dup function not working: %v", res)
  }
}

func TestStdlibStack7(t *testing.T) {
  tc := Init()
  tc = tc.Eval("42 , len")
  res := tc.GetAsString()
  if  res != "0" {
    t.Fatalf("in stack drop function not working: %v", res)
  }
}

func TestStdlibStack8(t *testing.T) {
  tc := Init()
  tc = tc.Eval("stack[1 , 42 ] len")
  res := tc.GetAsString()
  if  res != "1" {
    t.Fatalf("ina args drop function not working: %v", res)
  }
}
