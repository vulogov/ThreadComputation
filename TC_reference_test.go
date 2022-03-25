package ThreadComputation

import (
	"testing"
)

func TestReference1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("`stack[1 2 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestReference2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("`stack[42] !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("`stack[42] !: %v and shall be 42", res)
  }
}

func TestReference3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("`stack[] ![42]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("`stack[] ![42]: %v and shall be 42", res)
  }
}
