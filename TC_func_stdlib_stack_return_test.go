package ThreadComputation

import (
	"testing"
)

func TestStdlibStackReturn1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("1 2 3 | 42 _ <--")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
          t.Fatalf("1 2 3 | 42 _ <--: %v and shall be 42", res)
  }
}
