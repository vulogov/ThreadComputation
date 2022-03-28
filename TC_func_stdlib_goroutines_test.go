package ThreadComputation

import (
	"testing"
)

func TestGR1(t *testing.T) {
  tc := Init()
	tc = tc.Eval("wait")
  if tc.Errors() != 0 {
    t.Fatalf(tc.Error())
  }
}

func TestGR2(t *testing.T) {
  tc := Init()
	tc = tc.Eval("spawn[ 'Hello world from goroutine' print ] wait")
  if tc.Errors() != 0 {
    t.Fatalf(tc.Error())
  }
}

func TestGR3(t *testing.T) {
  // SetVariable("tc.Debuglevel", "debug")
  tc := Init()
	tc = tc.Eval("42 send recv")
  // SetVariable("tc.Debuglevel", "info")
  if tc.Errors() != 0 {
    t.Fatalf(tc.Error())
  }
  res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("42 send recv: %v", res)
  }
}
