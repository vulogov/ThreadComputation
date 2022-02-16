package ThreadComputation

import (
	"testing"
)


func TestUserFunc1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@SomeFun[]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestUserFunc2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@answer[42] answer")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@answer[42] answer  not working: %v", res)
	}
}

func TestUserFunc3(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@answer[\"Hello world!\"] answer")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "Hello world!" {
		t.Fatalf("@answer[\"Hello world!\"] answer  not working: %v", res)
	}
}

func TestUserFunc4(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@pi[3.14] pi")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3.14" {
		t.Fatalf("@pi[3.14] pi not working: %v", res)
	}
}

func TestUserFunc5(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@returntrue[#TRUE] returntrue")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("@returntrue[#TRUE] returntrue not working: %v", res)
	}
}

func TestUserFunc6(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@returnanswer[42 $answer] returnanswer answer ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@returnanswer[42 $answser] returnanswer not working: %v", res)
	}
}

func TestUserFunc7(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@returnanswer[stack[42]] returnanswer")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@returnanswer[stack[42]] returnanswer not working: %v", res)
	}
}

func TestUserFunc8(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@add[+] add[2 40]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@add[+] add[2 40] returnanswer not working: %v", res)
	}
}

func TestUserFunc10(t *testing.T) {
	tc := Init()
  tc = tc.Eval("@answer[ 41 (42) <- ] answer ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("@answer[ 41 (42) <- ] answer not working: %v", res)
	}
}
