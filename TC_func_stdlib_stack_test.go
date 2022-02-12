package ThreadComputation

import (
	"testing"
)


func TestStack1(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("1 2 3 | 42")
	if tc1.Errors() != 0 {
		t.Fatalf("1 2 3 | 42 parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("data in stack is not in order: %v and shall be 42", res)
	}
}

func TestStack2(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("1 2 3 (42 _)")
	if tc1.Errors() != 0 {
		t.Fatalf("1 2 3 (42 _) parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("1 2 3 (42 _) failed: %v and shall be 42", res)
	}
}

func TestStack3(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("1 2 3 (_[42])")
	if tc1.Errors() != 0 {
		t.Fatalf("1 2 3 (_[42]) parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("1 2 3 (_[42]) failed: %v and shall be 42", res)
	}
}

func TestStack4(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("1 2 3 | 42 | 4 5 <- ")
	if tc1.Errors() != 0 {
		t.Fatalf("1 2 3 | 42 | 4 5 <-  parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("1 2 3 | 42 | 4 5 <-  failed: %v and shall be 42", res)
	}
}

func TestStack5(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("42 | 1 2 3 | 4 5 -> ")
	if tc1.Errors() != 0 {
		t.Fatalf("42 | 1 2 3 | 4 5 ->  parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("42 | 1 2 3 | 4 5 ->  failed: %v and shall be 42", res)
	}
}

func TestStack6(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("1 42 2 << ")
	if tc1.Errors() != 0 {
		t.Fatalf("1 42 2 <<   parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("1 42 2 <<   failed: %v and shall be 42", res)
	}
}

func TestStack7(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("42 1 2 >> ")
	if tc1.Errors() != 0 {
		t.Fatalf("42 1 2 >>   parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("42 1 2 >>   failed: %v and shall be 42", res)
	}
}

func TestStack8(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("1 2 42 <<[2 1] ")
	if tc1.Errors() != 0 {
		t.Fatalf("1 2 42 <<[2 1]   parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("1 2 42 <<[2 1]   failed: %v and shall be 42", res)
	}
}

func TestStack9(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("-1 -2 -3 | 1 2 3 | 42 | 4 5 <-[5] ")
	if tc1.Errors() != 0 {
		t.Fatalf("-1 -2 -3 | 1 2 3 | 42 | 4 5 <-[5] parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("-1 -2 -3 | 1 2 3 | 42 | 4 5 <-[5]  failed: %v and shall be 42", res)
	}
}

func TestStack10(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("-1 -2 -3 | 1 2 3 | 42 | 4 5 ->[1 2] ")
	if tc1.Errors() != 0 {
		t.Fatalf("-1 -2 -3 | 1 2 3 | 42 | 4 5 ->[1 2] parse failed")
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("-1 -2 -3 | 1 2 3 | 42 | 4 5 ->[1 2]  failed: %v and shall be 42", res)
	}
}