package ThreadComputation

import (
	"testing"
)

func TestStdlibStackBackSend1(t *testing.T) {
	code := "1 2 3 ;[42]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
          t.Fatalf("%v failed: %v", code, res)
  }
}

func TestStdlibStackBackSend2(t *testing.T) {
	code := "1 2 ;[42] len"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "3" {
          t.Fatalf("%v failed: %v", code, res)
  }
}

func TestStdlibStackBackSend3(t *testing.T) {
	code := "stack [ ;[41] +[1] ]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
          t.Fatalf("%v failed: %v", code, res)
  }
}

func TestStdlibStackBackSend4(t *testing.T) {
	code := "local[let['p' 50]] stack [ ;[41] +[1] ] value"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
          t.Fatalf("%v failed: %v", code, res)
  }
}
