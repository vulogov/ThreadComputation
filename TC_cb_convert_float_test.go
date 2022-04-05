package ThreadComputation

import (
	"testing"
)

func Test_cb_float_convert1(t *testing.T) {
	code := "Float convert[Binary convert[3.14]] println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "3.14" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}
