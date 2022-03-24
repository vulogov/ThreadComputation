package ThreadComputation

import (
	"testing"
)

func TestStdlibSet1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("set[1 2 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
