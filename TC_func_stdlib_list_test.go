package ThreadComputation

import (
	"testing"
)

func TestStdlibList1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("list[1 2 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibList2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("list[1 'Hello' 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
