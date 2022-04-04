package ThreadComputation

import (
	"fmt"
	"testing"
)

func TestBinary1(t *testing.T) {
	code := "bin['Hello'] size"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "5" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestBinary2(t *testing.T) {
	code := "bin +['Hello'] size"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "5" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestBinary3(t *testing.T) {
	code := "bin['Hello'] =[bin['Hello']]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "true" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestBinary4(t *testing.T) {
	code := "bin['Hello'] =['Hello']"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "true" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestBinary5(t *testing.T) {
	code := "bin['Hello'] println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestBinary6(t *testing.T) {
	b := MakeBinary("12345")
	s := b.Slice(1,3)
	fmt.Println(s.String())
}
