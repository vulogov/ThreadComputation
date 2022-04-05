package ThreadComputation

import (
	"testing"
)

func TestStdlibSetLines1(t *testing.T) {
	code := "lines['Hello' 'Second line'] <-[pair[1 '42']] ->[1]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}
