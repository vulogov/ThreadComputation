package ThreadComputation

import (
	"testing"
)

func TestVars1(t *testing.T) {
	SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	tc = tc.Eval("42 $answer")
	SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestVars2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	tc = tc.Eval("stack [1 2 42 $answer]")
	// SetVariable("tc.Debuglevel", "info")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
//
// func TestVars3(t *testing.T) {
// 	// SetVariable("tc.Debuglevel", "debug")
// 	tc := Init()
// 	tc = tc.Eval("stack [1 2 42 $answer] answer")
// 	// SetVariable("tc.Debuglevel", "info")
// 	if tc.Errors() != 0 {
// 		t.Fatalf(tc.Error())
// 	}
// 	res := tc.GetAsString()
//   if res != "42" {
//     t.Fatalf("stack [1 2 42 $answer] answer failed: %v", res)
//   }
// }
//
// func TestVars4(t *testing.T) {
// 	// SetVariable("tc.Debuglevel", "debug")
// 	tc := Init()
// 	tc = tc.Eval(" 42 $answer answer")
// 	// SetVariable("tc.Debuglevel", "info")
// 	if tc.Errors() != 0 {
// 		t.Fatalf(tc.Error())
// 	}
// 	res := tc.GetAsString()
//   if res != "42" {
//     t.Fatalf("42 $answer answer failed: %v", res)
//   }
// }
