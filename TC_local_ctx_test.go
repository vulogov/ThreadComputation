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

func TestLocalCtx4(t *testing.T) {
	code := "context <-[let['answer' 42]] ! #answer"
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

func TestLocalCtx5(t *testing.T) {
	code := "context <-[let['answer' 42]] ->['answer']"
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

func TestLocalCtx6(t *testing.T) {
	code := "local[let['answer' 42]] local[let['true' true]] #answer"
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

func TestLocalCtx7(t *testing.T) {
	code := "local[let['true' true]] getcontext <-[let['answer' 42]] ! #answer"
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

func TestLocalCtx8(t *testing.T) {
	code := "local[let['value' 42]] □println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestLocalCtx9(t *testing.T) {
	code := "42 &  □println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
