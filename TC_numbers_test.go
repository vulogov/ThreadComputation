package ThreadComputation

import (
	"testing"
)

func TestNumbers1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("Numbers[1 Value[2 50.0] 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] +[1] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] numbers[1 2 3] ~* println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] =[numbers[1 2 3]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("numbers[1 2 3] =[numbers[1 2 3]] had failed: %v", res)
	}
}
