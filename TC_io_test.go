package ThreadComputation

import (
	"testing"
)

func TestIO1(t *testing.T) {
	code := "io['./examples/hello.txt'] println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestIO2(t *testing.T) {
	code := "io <-['./examples/hello.txt'] println"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestIO3(t *testing.T) {
	code := "Lines [io['./examples/hello.txt'] Read] size"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
  if res != "1" {
    t.Fatalf("%v is failed: %v", code, res)
  }
}

func TestIO4(t *testing.T) {
	code := "io['./examples/hello2.txt'] Write[io['./examples/hello.txt'] Read]]"
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval(code)
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
