package ThreadComputation

import (
	"testing"
)

func TestMatrix1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2] numbers[3 4]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestMatrix2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2 3] numbers[4 5]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestMatrix3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2 3] numbers[4 5]] size")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "6" {
    t.Fatalf("matrix[numbers[1 2 3] numbers[4 5]] size failed: %v", res)
  }
}

func TestMatrix4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2 3] numbers[4 5]] +[2] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestMatrix5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2 3] numbers[4 5]] +[matrix[numbers[1 2 3] numbers[4 5]]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestMatrix6(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2 3] numbers[4 5]] ^ println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestMatrix7(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[numbers[1 2 3] numbers[4 5]] matrix[numbers[1 2 3] numbers[4 5]] ~= ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "true" {
    t.Fatalf("matrix[numbers[1 2 3] numbers[4 5]] matrix[numbers[1 2 3] numbers[4 5]] ~= failed: %v", res)
  }
}

func TestMatrix8(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("matrix[list[numbers[1 2] numbers[3 4]]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
