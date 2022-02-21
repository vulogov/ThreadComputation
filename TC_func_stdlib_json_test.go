package ThreadComputation

import (
	"testing"
)

func TestJson1(t *testing.T) {
	tc := Init()
	tc = tc.Eval("json print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestJson2(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42 msg: \"Hello world!\" pi:3.14 truth:#TRUE} json print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestJson3(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42 msg: \"Hello world!\" pi:3.14 truth:#TRUE} json json print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestJson4(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42} json[read[\"@/examples/sample.json\"]] print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestJson5(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42} json[read[\"@/examples/sample.json\"]] +++[read[\"@/examples/sample3.json\"]] print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestJson6(t *testing.T) {
	tc := Init()
	tc = tc.Eval("{answer:42} json print answer")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("{answer:42} json print answer failed: %v", res)
	}
}
