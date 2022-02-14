package ThreadComputation

import (
	"github.com/gammazero/deque"
	"testing"
)

func testFunctionExists(l *TCExecListener, q *deque.Deque) (interface{}, error) {
	return 42, nil
}

func TestFunc1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("FunctionNotExists")
  if tc.Errors() == 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc11(t *testing.T) {
	tc := Init()
	tc.SetFunction("FunctionExists", testFunctionExists)
  tc = tc.Eval("FunctionExists")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc12(t *testing.T) {
	tc := Init()
	tc.SetFunction("FunctionExists", testFunctionExists)
  tc = tc.Eval("FunctionExists")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("stack[ stack[42]] not working: %v", res)
	}
}

func TestFunc2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("stack[ stack[42]]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("stack[ stack[42]] not working: %v", res)
	}
}
