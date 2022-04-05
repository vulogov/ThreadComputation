package ThreadComputation

import (
	"testing"
)

func TestStdlibLoop1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("3 loop[#0 println]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibLoop2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range[34 42 4] loop[#0 println]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlibLoop3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("range[34 42 4] loop[stack[#0]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("range[34 42 4] loop[stack[#0]] had failed: %v", res)
	}
}

func TestStdlibLoop4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("list[40 41 42] loop[stack[#0]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("list[40 41 42] loop[stack[#0]] had failed: %v", res)
	}
}

func TestStdlibLoop5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("list[40 41 42] ![pair['current' 1]] loop[stack[#0]] len")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
		t.Fatalf("list[40 41 42] ![pair['current' 1]] loop[stack[#0]] len had failed: %v", res)
	}
}

func TestStdlibLoop6(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2] numbers[3 4]] loop[stack[#0]] len")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
		t.Fatalf("matrix[numbers[1 2] numbers[3 4]] loop[stack[#0]] len had failed: %v", res)
	}
}

func TestStdlibLoop7(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2] numbers[3 4]] ![pair['current' 1]] loop[stack[#0]] len")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "1" {
		t.Fatalf("matrix[numbers[1 2] numbers[3 4]] ![pair['current' 1]] loop[stack[#0]] len had failed: %v", res)
	}
}

func TestStdlibLoop8(t *testing.T) {
	code := "bin['Hello'] loop[stack[#0]] len"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "5" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibLoop9(t *testing.T) {
	code := "bin['Hello'] ![let['buffer' 2]] loop[stack[#0]] len"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibLoop10(t *testing.T) {
	code := "'Hello' loop[stack[#0]] len"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "1" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibLoop11(t *testing.T) {
	code := "list[list[1 2 3] list[3 4 5]] loop[stack[#2]] ~+"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "6" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibLoop12(t *testing.T) {
	code := "list[dict['value' 41] dict['value' 1]] loop[stack[#value]] ~+"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibLoop13(t *testing.T) {
	code := "list[pair['value' 41] pair['value' 1]] loop[stack[#value]] ~+"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibLoop14(t *testing.T) {
	code := "list[pair['value' 41] pair['value' 1]] loop[stack[#2]] ~+"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}

func TestStdlibLoop15(t *testing.T) {
	code := "list[numbers[0 41 1] numbers[1 1 42]] loop[stack[#2]] ~+"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}
