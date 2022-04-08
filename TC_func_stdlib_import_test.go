package ThreadComputation

import (
	"testing"
)

func TestStdlibImport1(t *testing.T) {
	code := "'Hello' export import"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "Hello" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibImport2(t *testing.T) {
	code := "42 export import"
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
