package ThreadComputation

import (
	"testing"
)

func TestCond1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 false ?41")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("Not pushed to stack: %v", res)
  }
}

func TestCond2(t *testing.T) {
	SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("41 true ?42")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("Not pushed to stack: %v", res)
  }
}
