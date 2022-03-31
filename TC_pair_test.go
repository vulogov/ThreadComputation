package ThreadComputation

import (
	"testing"
)

func TestPair1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("pair[1 2] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
