package ThreadComputation

import (
	"testing"
)

func TestStdlibMathDec1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("43 --")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 43 -- had failed: %v", res)
	}
}

func TestStdlibMathDec2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2] numbers[3 4]] -- println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibMathDec3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] -- println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
