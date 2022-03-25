package ThreadComputation

import (
	"testing"
)

func TestStdlibOnError1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("throw[\"test\"]")
	if tc.Errors() == 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibOnError2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("errors.Handle throw[\"test\"] onError[errors.Clear] errors.Check")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "false" {
          t.Fatalf("errors.Handle throw[\"test\"] onError[errors.Clear] errors.Check  shall be false and it is %v", res)
  }
}
