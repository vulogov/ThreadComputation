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
