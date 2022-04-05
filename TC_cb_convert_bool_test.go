package ThreadComputation

import (
	"testing"
)

func Test_cb_bool_convert1(t *testing.T) {
	code := "Bool convert[Binary convert[true]] println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "true" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}
