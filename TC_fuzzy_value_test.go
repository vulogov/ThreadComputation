package ThreadComputation

import (
	"testing"
)

func TestFuzzyValue1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	SetVariable("tc.Debuglevel", "info")
	// tc = tc.Eval("{42 90}")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFuzzyValue2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("{41 90} {1} ~+ value")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("{41 90} {42} ~+ value had failed: %v", res)
	}
}

func TestFuzzyValue3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("{stack[40 41 42] 90} len")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3" {
		t.Fatalf("{stack[40 41 42] 90} len had failed: %v", res)
	}
}
