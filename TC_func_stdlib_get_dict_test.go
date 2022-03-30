package ThreadComputation

import (
	"testing"
)

func TestStdlibGetDict1(t *testing.T) {
	SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("dict['answer' 42] ->['answer']")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("dict['answer' 42] ->['answer'] had failed: %v", res)
	}
}
