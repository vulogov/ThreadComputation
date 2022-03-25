package ThreadComputation

import (
	"testing"
)

func TestStdliblambda1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("lambda[42] !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("lambda[42] ! had failed: %v", res)
	}
}

func TestStdliblambda2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("λ[42] !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("λ[42] ! had failed: %v", res)
	}
}

func TestStdliblambda3(t *testing.T) {
	SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("λ[#0] ![42]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("λ[#0] ![42] had failed: %v", res)
	}
}
