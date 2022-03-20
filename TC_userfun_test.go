package ThreadComputation

import (
	"testing"
)

func TestUserfun1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	tc = tc.Eval("@answer[42] answer")
	// SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@answer[42] answer had failed: %v", res)
	}
}

func TestUserfun2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	tc = tc.Eval("@answer[42] stack[ answer ]")
	// SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@answer[42] stack[ answer ] had failed: %v", res)
	}
}
