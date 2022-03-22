package ThreadComputation

import (
	"testing"
)

func TestNs1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 [hello: 1 2 3 ;;")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("42 [hello: 1 2 3 ;;  had failed: %v", res)
	}
}
