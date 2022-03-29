package ThreadComputation

import (
	"testing"
)

func TestStdlibMath1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("40 +[2]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 40 +[2] had failed: %v", res)
	}
}

func TestStdlibMath2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("2 40 ~+")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 2 40 ~+ had failed: %v", res)
	}
}

func TestStdlibMath3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("44 -[2]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 44 -[2] had failed: %v", res)
	}
}

func TestStdlibMath4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("84 /[2]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 84 /[2] had failed: %v", res)
	}
}

func TestStdlibMath5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("21 *[2]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 21 *[2] had failed: %v", res)
	}
}

func TestStdlibMath6(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("'4' +[2]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" '4' +[2] had failed: %v", res)
	}
}

func TestStdlibMath7(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("'4' *[2]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "44" {
		t.Fatalf(" '4' *[2] had failed: %v", res)
	}
}

func TestStdlibMath8(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("1 /[0] println")
	if tc.Errors() == 0 {
		t.Fatalf(tc.Error())
	}
}
