package ThreadComputation

import (
	"testing"
)

func TestNone1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("none println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNone2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("None println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
