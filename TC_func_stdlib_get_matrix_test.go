package ThreadComputation

import (
	"testing"
)

func TestStdlibGetMatrix1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2] numbers[3 4]] ->[pair[1 1]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4" {
		t.Fatalf("matrix[numbers[1 2] numbers[3 4]] ->[pair[1 1]] had failed: %v", res)
	}
}
