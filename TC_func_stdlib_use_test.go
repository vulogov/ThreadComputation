package ThreadComputation

import (
	"testing"
)



func TestUse1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("use[\"https://raw.githubusercontent.com/vulogov/ThreadComputation/main/examples/answer.tc\"]")
  if tc.Errors() != 0 {
    t.Fatalf(tc.Error())
  }
  res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("use[\"https://raw.githubusercontent.com/vulogov/ThreadComputation/main/examples/answer.tc\"] failed: %v", res)
  }
}

func TestUse2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("\"https://raw.githubusercontent.com/vulogov/ThreadComputation/main/examples/answer.tc\" use")
  if tc.Errors() != 0 {
    t.Fatalf(tc.Error())
  }
  res := tc.GetAsString()
  if res != "42" {
    t.Fatalf("\"https://raw.githubusercontent.com/vulogov/ThreadComputation/main/examples/answer.tc\" use failed: %v", res)
  }
}
