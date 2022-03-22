package ThreadComputation

import (
	"testing"
)

func TestPos1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("@echo[#0] echo[42]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@echo[#0] echo[42]  had failed: %v", res)
	}
}
