package ThreadComputation

import (
	"testing"
)


func TestStdlib1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("\"Hello world\" print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib3(t *testing.T) {
	tc := Init()
  tc = tc.Eval("3.14 print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib4(t *testing.T) {
	tc := Init()
  tc = tc.Eval("#TRUE print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
