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

func TestFunc5(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 #FALSE ?stack[ 41 ]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 42 #FALSE ?stack[ 41 ] not working: %v", res)
	}
}

func TestFunc51(t *testing.T) {
	tc := Init()
  tc = tc.Eval("41 #TRUE ?stack[ 42 ]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 41 #TRUE ?stack[ 42 ] not working: %v", res)
	}
}

func TestFunc6(t *testing.T) {
	tc := Init()
  tc = tc.Eval("1 1 2 2 3 3 set print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc7(t *testing.T) {
	tc := Init()
  tc = tc.Eval("set[42] unset ")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" set[42] unset not working: %v", res)
	}
}

func TestFunc8(t *testing.T) {
	tc := Init()
  tc = tc.Eval("set +++[42] unset")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" set +++[42] unset not working: %v", res)
	}
}

func TestFunc9(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 set.New +++ unset")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 42 set.New +++ unset not working: %v", res)
	}
}

func TestFunc10(t *testing.T) {
	tc := Init()
  tc = tc.Eval("set.New[1 2] +++[set.New[3]] print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc110(t *testing.T) {
	tc := Init()
  tc = tc.Eval("set.New[1 2] +++[set.New[3]] set.Len")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3" {
		t.Fatalf(" set.New[1 2] +++[set.New[3]] set.Len not working: %v", res)
	}
}

func TestFunc111(t *testing.T) {
	tc := Init()
  tc = tc.Eval("set.New[1 2] +++[set.New[3]] unset len")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3" {
		t.Fatalf(" set.New[1 2] +++[set.New[3]] unset len not working: %v", res)
	}
}

func TestFunc200(t *testing.T) {
	tc := Init()
  tc = tc.Eval("set.New[1 42] ---[1] unset")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("set.New[1 42] ---[1] unset not working: %v", res)
	}
}

func TestFunc201(t *testing.T) {
	tc := Init()
  tc = tc.Eval("set.New[42] set.New[1 42] --- unset")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("set.New[42] set.New[1 42] --- unset not working: %v", res)
	}
}

func TestFunc202(t *testing.T) {
	tc := Init()
  tc = tc.Eval("41 in[42]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "false" {
		t.Fatalf("41 in[42] not working: %v", res)
	}
}

func TestFunc203(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 set in[42]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("42 set in[42] not working: %v", res)
	}
}

func TestFunc204(t *testing.T) {
	tc := Init()
  tc = tc.Eval("&")
  if tc.Errors() == 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc205(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@test[ 1 + dup 5 > printStack ?test] 0 test")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestFunc206(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@test[ ++[1] dup 5 > ?test] 0 test ")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "5" {
		t.Fatalf("@test[ ++[1] dup 5 > ?test] 0 test not working: %v", res)
	}
}
