package ThreadComputation

import (
	"testing"
)

func TestGR1(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("wait")
	if tc1.Errors() != 0 {
		t.Fatalf("wait parse failed")
	}
}

func TestGR2(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("spawn\\ 'Hello world from goroutine' print \\ wait")
	if tc1.Errors() != 0 {
		t.Fatalf("spawn\\ 'Hello world' print \\ wait parse failed")
	}
}

func TestGR3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
  tc1 := tc.Eval("42 send recv")
	// SetVariable("tc.Debuglevel", "info")
	if tc1.Errors() != 0 {
		t.Fatalf("42 send recv parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("42 send recv failed: %v", res)
	}
}

func TestGR4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
  tc1 := tc.Eval(":test(42 send) recv['test']")
	// SetVariable("tc.Debuglevel", "info")
	if tc1.Errors() != 0 {
		t.Fatalf(":test(42 send) recv['test'] parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(":test(42 send) recv['test'] failed: %v", res)
	}
}

func TestGR5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
  tc1 := tc.Eval(":test() 42 send['test'] S['test'] recv")
	// SetVariable("tc.Debuglevel", "info")
	if tc1.Errors() != 0 {
		t.Fatalf(":test() 42 send['test'] S['test'] recv parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(":test() 42 send['test'] S['test'] recv failed: %v", res)
	}
}
