package ThreadComputation

import (
	"testing"
)

func TestRange1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestRange2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range[4] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestRange3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range[40 42] loop[#0 println]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestRange4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range[34 42 4] loop[#0 println]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
