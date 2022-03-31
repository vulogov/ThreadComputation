package ThreadComputation

import (
	"testing"
)

func TestDict1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("dict['answer' 42] println ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
