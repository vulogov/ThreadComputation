package ThreadComputation

import (
	"testing"
)

func TestStdlibSetJson1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("json['answer' 42] <-['{\"true\":true}'] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibSetJson2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("json['answer' 42] <-[json['true' true]] ->['true']")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "true" {
    t.Fatalf("json['answer' 42] <-[json['true' true]] ->['true']: %v and shall be true", res)
  }
}

func TestStdlibSetJson3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("json['answer' 42] <-[dict['true' true]] ->['true']")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "true" {
    t.Fatalf("json['answer' 42] <-[dict['true' true]] ->['true']: %v and shall be true", res)
  }
}
