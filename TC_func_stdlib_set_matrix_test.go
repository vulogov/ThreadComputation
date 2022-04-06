package ThreadComputation

import (
	"testing"
)

func TestStdlibSetMatrix1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2] numbers[3 4]] <-[pair[pair[1 1] 42]] ->[pair[1 1]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("matrix[numbers[1 2] numbers[3 4]] <-[pair[pair[1 1] 42]] ->[pair[1 1]]: %v and shall be 42", res)
  }
}
