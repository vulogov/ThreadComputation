package ThreadComputation

import (
	"testing"
)

func TestCode1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("code[41 1 ~+ ] println ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}

}
