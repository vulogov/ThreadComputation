package ThreadComputation

import (
	"testing"
)



func TestStdlibSampleDataframe1(t *testing.T) {
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
	sample["Answers"] size
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
		t.Fatalf("%v had failed: %v", code, res)
	}
}
