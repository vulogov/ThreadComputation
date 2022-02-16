package ThreadComputation

import (
	"testing"
)



func TestCmp1(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("42 42 =")
	if tc1.Errors() != 0 {
		t.Fatalf("42 42 = parse failed")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("42 42 = failed: %v", res)
	}
}

func TestCmp2(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("42.0  =[42.0]")
	if tc1.Errors() != 0 {
		t.Fatalf("42.0 =[42.0] parse failed")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("42.0 =[42.0] failed: %v", res)
	}
}

func TestCmp3(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval(">[42 41]")
	if tc1.Errors() != 0 {
		t.Fatalf(">[42 41] parse failed")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf(">[42 41] failed: %v", res)
	}
}

func TestCmp31(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval(">=[42 41]")
	if tc1.Errors() != 0 {
		t.Fatalf(">=[42 41] parse failed")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf(">=[42 41] failed: %v", res)
	}
}

func TestCmp4(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("<>[\"Hello\" \"World\"]")
	if tc1.Errors() != 0 {
		t.Fatalf("<>[\"Hello\" \"World\"] parse failed")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("<>[\"Hello\" \"World\"] failed: %v", res)
	}
}

func TestCmp5(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("<[41 42]")
	if tc1.Errors() != 0 {
		t.Fatalf("<[41 42] parse failed")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("<[41 42] failed: %v", res)
	}
}

func TestCmp51(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("<=[41 42]")
	if tc1.Errors() != 0 {
		t.Fatalf("<=[41 42] parse failed")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("<=[41 42] failed: %v", res)
	}
}

func TestCmp6(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("#FALSE #FALSE and")
	if tc1.Errors() != 0 {
		t.Fatalf("#FALSE #FALSE and")
	}
	res := tc.GetAsString()
	if res != "false" {
		t.Fatalf("#FALSE #FALSE and failed: %v", res)
	}
}

func TestCmp7(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("#FALSE or[#TRUE]")
	if tc1.Errors() != 0 {
		t.Fatalf("#FALSE or[#TRUE]")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("#FALSE or[#TRUE] failed: %v", res)
	}
}

func TestCmp8(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("or[#FALSE #FALSE]")
	if tc1.Errors() != 0 {
		t.Fatalf("or[#FALSE #FALSE]")
	}
	res := tc.GetAsString()
	if res != "false" {
		t.Fatalf("or[#FALSE #FALSE] failed: %v", res)
	}
}

func TestCmp9(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("#FALSE not")
	if tc1.Errors() != 0 {
		t.Fatalf("#FALSE not")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("#FALSE not failed: %v", res)
	}
}

func TestCmp10(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("¬[#FALSE]")
	if tc1.Errors() != 0 {
		t.Fatalf("¬[#FALSE]")
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("¬[#FALSE] failed: %v", res)
	}
}
