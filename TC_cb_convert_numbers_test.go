package ThreadComputation

import (
	"testing"
)

func Test_cb_numbers_convert1(t *testing.T) {
	code := "List convert[numbers[1 2 3]] size"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "3" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}
