package ThreadComputation

import (
	"testing"
)

func TestExported1(t *testing.T) {
	code := "Exported convert[Binary['Hello']] size"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "16" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestExported2(t *testing.T) {
	code := "Int convert[ Binary convert[Exported convert[Binary[42]]]]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}
