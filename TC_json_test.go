package ThreadComputation

import (
	"testing"
)

func TestJson1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("json['answer' 42] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestJson2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("jsonUniq println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
