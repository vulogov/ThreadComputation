package ThreadComputation

import (
	"testing"
)

func TestStdlibMathInc1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("41 ++")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 40 ++ had failed: %v", res)
	}
}

func TestStdlibMathInc2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2] numbers[3 4]] ++ println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibMathInc3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] ++ println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
