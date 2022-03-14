package ThreadComputation

import (
	"testing"
)

func TestStdlibStr1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("stack[\"42=%v\" str.format[42]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibString1(t *testing.T) {
  tc := Init()
	tc1 := tc.Eval("\"Hello\" str.upper")
  if tc1.Errors() != 0 {
    t.Fatalf("\"Hello\" str.upper parse failed")
  }
  res := tc.GetAsString()
  if res != "HELLO" {
    t.Fatalf("\"Hello\" str.upper parse failed: %v", res)
  }
}

func TestStdlibString2(t *testing.T) {
  tc := Init()
	tc1 := tc.Eval("\"HelLo\" str.lower")
  if tc1.Errors() != 0 {
    t.Fatalf("\"HelLo\" str.lower parse failed")
  }
  res := tc.GetAsString()
  if res != "hello" {
    t.Fatalf("\"HelLo\" str.lower parse failed: %v", res)
  }
}

func TestStdlibString3(t *testing.T) {
  tc := Init()
	tc1 := tc.Eval("\"helLo\" str.lower str.title")
  if tc1.Errors() != 0 {
    t.Fatalf("\"helLo\" str.lower str.title parse failed")
  }
  res := tc.GetAsString()
  if res != "Hello" {
    t.Fatalf("\"helLo\" str.lower str.title failed: %v", res)
  }
}
