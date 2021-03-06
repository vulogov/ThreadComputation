package ThreadComputation

import (
	"testing"
)

func TestDataframe1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("data[] println ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestDataframe2(t *testing.T) {
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
	println
	`
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestDataframe3(t *testing.T) {
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
	println
	`
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestDataframe4(t *testing.T) {
	code := `
	List
	convert[
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
	]
	println
	`
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
