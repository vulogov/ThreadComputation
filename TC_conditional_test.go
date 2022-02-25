package ThreadComputation

import (
	"testing"
)

func TestCond1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
  tc = tc.Eval("0 false\\ +[1] dup print 10 > \\ ")
	// SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "10" {
		t.Fatalf("0 false\\ +[1] dup print 10 > \\  not working: %v", res)
	}
}

func TestCond2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
  tc = tc.Eval("10 true\\ -[1] dup print 0 >= \\ ")
	// SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "0" {
		t.Fatalf("10 true\\ -[1] dup print 0 >= \\  not working: %v", res)
	}
}
