package ThreadComputation

import (
	"testing"
)

func TestFilter1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
  tc = tc.Eval("41 42 filter\\ 42 = \\ ")
	// SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("41 42 filter\\ 42 = \\  not working: %v", res)
	}
}

func TestFilter2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
  tc = tc.Eval("41 42 filter\\ 42 = \\ len")
	// SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "1" {
		t.Fatalf("41 42 filter\\ 42 = \\  len not working: %v", res)
	}
}
