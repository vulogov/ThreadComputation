package ThreadComputation

import (
	"testing"
)

func TestFuzzy1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("Just[42] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFuzzy2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("Never[41] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFuzzy3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("Likely[41.5] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFuzzy4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("Value[41 30.0] +[Value[1 10.0]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFuzzy5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("41 +[Value[1 10.0]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFuzzy6(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 0 1 ~^+ println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFuzzy7(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("Value[41 70.0] Value[3 10.0] Value[1 90.0] ~^+ println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
