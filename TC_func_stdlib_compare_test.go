package ThreadComputation

import (
	"testing"
)

func TestStdlibCompare1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 =[42]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("42 =[42] had failed: %v", res)
	}
}

func TestStdlibCompare2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 42 ~=")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("42 42 ~= had failed: %v", res)
	}
}

func TestStdlibCompare3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 Value[42 50.0] ~= value")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("42 Value[42 50.0] ~= value had failed: %v", res)
	}
}

func TestStdlibCompare4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 Value[41 50.0] ~< value")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "false" {
		t.Fatalf("42 Value[41 50.0] ~< value had failed: %v", res)
	}
}
