package ThreadComputation

import (
	"testing"
)

func TestStdlibLoop1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("3 loop[#0 println]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibLoop2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range[34 42 4] loop[#0 println]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibLoop3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range[34 42 4] loop[stack[#0]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("range[34 42 4] loop[stack[#0]] had failed: %v", res)
	}
}
