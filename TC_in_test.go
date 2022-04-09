package ThreadComputation

import (
	"testing"
)

func TestIn1(t *testing.T) {
	code := "global[let['inContext' false]] in[] #inContext"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "false" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestIn2(t *testing.T) {
	code := "global[let['inContext' false]] in[stack[#inContext]]"
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

func TestIn3(t *testing.T) {
	code := "local[let['answer' 42]] in[here['answser' 41]] #answer"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestIn4(t *testing.T) {
	code := "local[let['answer' 41]] in[here['answer' 42] stack[#answer]]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}
