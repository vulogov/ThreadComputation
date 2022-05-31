package ThreadComputation

import (
	"testing"
)

func TestStdlibCarcdr1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("list[42 43 44] car")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("list[42 43 44] car had failed: %v", res)
	}
}

func TestStdlibCarcdr2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("list[42 43 44] cdr size")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
		t.Fatalf("list[42 43 44] cdr size had failed: %v", res)
	}
}

func TestStdlibCarcdr3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[42 43 44] car")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("numbers[42 43 44] car had failed: %v", res)
	}
}

func TestStdlibCarcdr4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[42 43 44] cdr size")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
		t.Fatalf("numbers[42 43 44] cdr size had failed: %v", res)
	}
}

func TestStdlibCarcdr5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("pair[1 2] car")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "1" {
		t.Fatalf("pair[1 2] car had failed: %v", res)
	}
}

func TestStdlibCarcdr6(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("pair[1 2] cdr")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
		t.Fatalf("pair[1 2] cdr had failed: %v", res)
	}
}

func TestStdlibCarcdr7(t *testing.T) {
	code := `
	data[
		pair[
			"Answers"
			Int
		]
		pair[
			"Comment"
			String
		]
	]
	+[
	  row[
			"Answers"
			42
			"Comment"
			"This is the Answer"
		]
		row[
			"Answers"
			41
			"Comment"
			"This is not an answer"
		]
	 ]
	car sample["Answers"] ->[0]
	`
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

func TestStdlibCarcdr8(t *testing.T) {
	code := `
	data[
		pair[
			"Answers"
			Int
		]
		pair[
			"Comment"
			String
		]
	]
	+[
	  row[
			"Answers"
			42
			"Comment"
			"This is the Answer"
		]
		row[
			"Answers"
			41
			"Comment"
			"This is not an answer"
		]
	 ]
	cdr sample["Answers"] ->[0]
	`
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "41" {
		t.Fatalf("%v had failed: %v", code, res)
	}
}
