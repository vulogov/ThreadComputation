package ThreadComputation

import (
	"testing"
)

func TestDmap1(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{}")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestDmap2(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42}")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestDmap3(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42} answer")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("{answer:42} answer failed: %v", res)
	}
}

func TestDmap4(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42 msg: \"Hello world!\" pi:3.14 truth:#TRUE} print ")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}
