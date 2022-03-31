package ThreadComputation

import (
	"testing"
)

func TestStdlibGetJson1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("json['answer' 42] ->['answer']")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("json['answer' 42] ->['answer'] had failed: %v", res)
	}
}

func TestStdlibGetJson2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("json['answer' 42] ->['not_answer']")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.Get()
	if ! IsNone(res) {
		t.Fatalf("json['answer' 42] ->['not_answer'] had failed: %v", res)
	}
}
