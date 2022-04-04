package ThreadComputation

import (
	"testing"
)

func TestLines1(t *testing.T) {
	code := "lines['Hello' 'Second line'] size"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "2" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestLines2(t *testing.T) {
	code := "lines['Hello' 'Second line'] +['Third line'] size"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "3" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestLines3(t *testing.T) {
	code := "lines['Hello' 'Second line'] ! loop[#0 println]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestLines4(t *testing.T) {
	code := "lines['Hello' 'Second line'] loop[stack[#0]] len"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}
