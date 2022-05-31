package ThreadComputation

import (
	"testing"
)

func TestStdlibSize1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("42 size")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "1" {
          t.Fatalf("42 size  shall be 1 and it is %v", res)
  }
}

func TestStdlibSize2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("list[1 2 3] size")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "3" {
          t.Fatalf("list[1 2 3] size  shall be 3 and it is %v", res)
  }
}

func TestStdlibSize3(t *testing.T) {
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
	 ]
	size
	`
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "1" {
      t.Fatalf("size of the dataframe is %v", res)
  }
}
