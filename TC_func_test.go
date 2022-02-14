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

func TestFunc13(t *testing.T) {
	tc := Init()
	tc2 := Init()
	tc.SetFunction("FunctionNotExists", testFunctionExists)
  tc = tc2.Eval("FunctionNotExists")
  if tc.Errors() == 0 {
		t.Fatalf(tc.Error())
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

func TestFunc3(t *testing.T) {
	tc := Init()
  tc = tc.Eval("`print[42] !")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc31(t *testing.T) {
	tc := Init()
  tc = tc.Eval("`stack[42] !")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("stack[ 42 ] ! not working: %v", res)
	}
}

func TestFunc32(t *testing.T) {
	tc := Init()
  tc = tc.Eval("`stack ![42]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("stack ![42] not working: %v", res)
	}
}

func TestFunc4(t *testing.T) {
	tc := Init()
  tc = tc.Eval("`+[40 2] !*")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("`+[40 2] !* not working: %v", res)
	}
}

func TestFunc41(t *testing.T) {
	tc := Init()
  tc = tc.Eval("`+[40 2] !*")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" !*[`+[40 2]] not working: %v", res)
	}
}
