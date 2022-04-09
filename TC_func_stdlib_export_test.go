package ThreadComputation

import (
	"testing"
)

func TestStdlibExport1(t *testing.T) {
	code := "'Hello' export size"
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
