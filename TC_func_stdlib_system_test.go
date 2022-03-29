package ThreadComputation

import (
	"testing"
)

func TestSystem1(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("[hello: name _ ;;")
	if tc1.Errors() != 0 {
		t.Fatalf("[hello: name _ ;; parse failed")
	}
	res := tc.GetAsString()
	if res != "hello" {
		t.Fatalf("[hello: name _ ;; had failed: %v", res)
	}
}

func TestSystem2(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("[hello: ;; previous")
	if tc1.Errors() != 0 {
		t.Fatalf("[hello: ;; previous parse failed")
	}
	res := tc.GetAsString()
	if res != "hello" {
		t.Fatalf("[[hello: ;; previous had failed: %v", res)
	}
}
