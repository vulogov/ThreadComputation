package ThreadComputation

import (
	"testing"
)

func TestStdlibConvert1(t *testing.T) {
	SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("Int convert['42']")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("Int convert['42'] had failed: %v", res)
	}
}
