package ThreadComputation

import (
	"testing"
)


func TestLambda1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("lambda\\42\\ print")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestLambda2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("lambda\\42\\ !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("lambda\\42\\ ! not working: %v", res)
	}
}
