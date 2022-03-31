package ThreadComputation

import (
	"strconv"
	"testing"
)



func TestNeural1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("neural println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNeural2(t *testing.T) {
	code1 := `
	/* Run Neutral XOR sample */
	neural[
		pair[
			numbers[0 0]
			numbers[0]
		]
		pair[
			numbers[0 1]
			numbers[1]
		]
		pair[
			numbers[1 0]
			numbers[1]
		]
		pair[
			numbers[1 1]
			numbers[0]
		]
	]
	train
	->[
		numbers[1 0]
	]
	car
	`
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code1)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res, _ := strconv.ParseFloat(tc.GetAsString(), 64)
	if res < 0.9 {
		t.Fatalf("Neural net had failed  had failed: %v", res)
	}
}
