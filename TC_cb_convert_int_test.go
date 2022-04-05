package ThreadComputation

import (
	"testing"
)

func Test_cb_int_convert1(t *testing.T) {
	code := "Int convert[Binary convert[42]] println"
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
