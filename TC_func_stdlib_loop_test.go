package ThreadComputation

import (
	"testing"
)

func TestStdlibLoop1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	SetVariable("tc.Debuglevel", "info")
	// tc = tc.Eval("3 loop[#0 println]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
