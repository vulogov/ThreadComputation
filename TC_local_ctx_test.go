package ThreadComputation

import (
	"testing"
)

func TestLocalCtx1(t *testing.T) {
	code := "local[let['answer' 42]] #answer"
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

func TestLocalCtx2(t *testing.T) {
	code := "context[let['answer' 42]] println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestLocalCtx3(t *testing.T) {
	code := "context ![let['answer' 42]] #answer"
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
