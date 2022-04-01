package ThreadComputation

import (
	"fmt"
	"testing"
)

func TestNumbers1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers2(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("JustNumbers[1 Value[2 50.0] 3] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers3(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] +[1] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers4(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] numbers[1 2 3] ~* println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers5(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[1 2 3] =[numbers[1 2 3]]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "true" {
		t.Fatalf("numbers[1 2 3] =[numbers[1 2 3]] had failed: %v", res)
	}
}

func TestNumbers6(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("numbers[list[1 2 3]] println")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestNumbers7(t *testing.T) {
	n := MakeNumbers()
	n.Add(float64(1))
	n.Add(float64(2))
	n.Add(float64(3))
	fmt.Printf("%v\n", n.String())
	n.RotateLeft()
	fmt.Printf("%v\n", n.String())
}

func TestNumbers8(t *testing.T) {
	n := MakeNumbers()
	n.Add(float64(1))
	n.Add(float64(2))
	n.Add(float64(3))
	fmt.Printf("%v\n", n.String())
	n.RotateRight()
	fmt.Printf("%v\n", n.String())
}
