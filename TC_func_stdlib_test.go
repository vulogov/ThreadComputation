package ThreadComputation

import (
	"testing"
)


func TestStdlib1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib11(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("Not pushed to stack: %v", res)
	}
}

func TestStdlib2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("\"Hello world\" print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib3(t *testing.T) {
	tc := Init()
  tc = tc.Eval("3.14 print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib4(t *testing.T) {
	tc := Init()
  tc = tc.Eval("#TRUE print")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib5(t *testing.T) {
	tc := Init()
  tc = tc.Eval("print[1 2 3]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib6(t *testing.T) {
	tc := Init()
  tc = tc.Eval("1 2 3 printAll")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib7(t *testing.T) {
	tc := Init()
  tc = tc.Eval("stack[1 2 3] printAll")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib8(t *testing.T) {
	tc := Init()
  tc = tc.Eval("stack[1 2 3] len")
	res := tc.GetAsString()
  if res != "3" {
		t.Fatalf("len function not working: %v", res)
	}
}

func TestStdlib9(t *testing.T) {
	tc := Init()
  tc = tc.Eval("stack[1 2 3] , len")
  if tc.GetAsString() != "2" {
		t.Fatalf("drop function not working")
	}
}

func TestStdlib10(t *testing.T) {
	tc := Init()
  tc = tc.Eval("stack[1 2 3] dup len")
	res := tc.GetAsString()
  if res != "4" {
		t.Fatalf("dup function not working: %v", res)
	}
}

func TestStdlib12(t *testing.T) {
	tc := Init()
  tc = tc.Eval("stack[1 2 3] clr len")
	res := tc.GetAsString()
  if  res != "0" {
		t.Fatalf("clr function not working: %v", res)
	}
}

func TestStdlib13(t *testing.T) {
	tc := Init()
  tc = tc.Eval("`print print")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib131(t *testing.T) {
	tc := Init()
  tc = tc.Eval("`print[42 41 40 39] print")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
}

func TestStdlib14(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" 2 40 `+ attr !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("2 40 `+ attr ! had failed: %v", res)
	}
}

func TestStdlib141(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" 2 44 `- attr !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("2 44 `- attr ! had failed: %v", res)
	}
}

func TestStdlib1412(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" `- attr[44 2] !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("`- attr[44 2] ! had failed: %v", res)
	}
}

func TestStdlibMath1(t *testing.T) {
	tc := Init()
  tc = tc.Eval("+[2 2]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4" {
		t.Fatalf("+[2 2] had failed: %v", res)
	}
}

func TestStdlibMath2(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 2 +")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4" {
		t.Fatalf("2 2 + had failed: %v", res)
	}
}

func TestStdlibMath3(t *testing.T) {
	tc := Init()
  tc = tc.Eval("-[43 1]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("-[43 1] had failed: %v", res)
	}
}

func TestStdlibMath4(t *testing.T) {
	tc := Init()
  tc = tc.Eval("3 3 3 *")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "27" {
		t.Fatalf("3 3 3 * had failed: %v", res)
	}
}

func TestStdlibMath5(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 9 /")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4.5" {
		t.Fatalf("2 9 / had failed: %v", res)
	}
}

func TestStdlibMath6(t *testing.T) {
	tc := Init()
  tc = tc.Eval("/[9 2]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4.5" {
		t.Fatalf("/[9 2] had failed: %v", res)
	}
}

func TestStdlibMath61(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 9 /")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4.5" {
		t.Fatalf("2 9 / had failed: %v", res)
	}
}


func TestStdlibMath7(t *testing.T) {
	tc := Init()
  tc = tc.Eval("max[9 2]")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "9" {
		t.Fatalf("max[9 2] had failed: %v", res)
	}
}

func TestStdlibMath8(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 9 min")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
		t.Fatalf("2 9 min had failed: %v", res)
	}
}

func TestStdlibMath9(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 9 mean")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "5.5" {
		t.Fatalf("2 9 mean had failed: %v", res)
	}
}

func TestStdlibMath10(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 9 variance")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "24.5" {
		t.Fatalf("2 9 variance had failed: %v", res)
	}
}

func TestStdlibMath11(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 10 12 skew")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "-1.4578629673213053" {
		t.Fatalf("2 10 12 skew had failed: %v", res)
	}
}

func TestStdlibMath12(t *testing.T) {
	tc := Init()
  tc = tc.Eval("2 10 12 deviation")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "5.291502622129181" {
		t.Fatalf("2 10 12 deviation had failed: %v", res)
	}
}

func TestStdlibMath13(t *testing.T) {
	tc := Init()
  tc = tc.Eval("-42 abs")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("-42 abs had failed: %v", res)
	}
}

func TestStdlibMath14(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 exp")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "1.7392749415205012e+18" {
		t.Fatalf("42 exp had failed: %v", res)
	}
}

func TestStdlibMath15(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 exp2")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "4.398046511104e+12" {
		t.Fatalf("42 exp2 had failed: %v", res)
	}
}

func TestStdlibMath16(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 log")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3.7376696182833684" {
		t.Fatalf("42 log had failed: %v", res)
	}
}

func TestStdlibMath17(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 log10")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "1.6232492903979006" {
		t.Fatalf("42 log10 had failed: %v", res)
	}
}

func TestStdlibMath18(t *testing.T) {
	tc := Init()
  tc = tc.Eval("42 sqrt")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "6.48074069840786" {
		t.Fatalf("42 sqrt had failed: %v", res)
	}
}

func TestStdlibMath19(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" 2 2 + 6 -")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "2" {
		t.Fatalf("2 2 + 6 - had failed: %v", res)
	}
}

func TestStdlibMath20(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" TheAnswer ")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("TheAnswer: %v", res)
	}
}

func TestStdlibMath21(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" TheAnswer[3.14] ")
  if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "3.14" {
		t.Fatalf("TheAnswer[3.14]: %v", res)
	}
}

func TestStdlib15(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" \"2 40 +\" !")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" \"2 40 +\" ! had failed: %v", res)
	}
}

func TestStdlib16(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" !*[\"2 40 +\"]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" !*[\"2 40 +\"] had failed: %v", res)
	}
}

func TestStdlib17(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" 42 ! ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf(" 42 ! had failed: %v", res)
	}
}

func TestStdlib18(t *testing.T) {
	tc := Init()
  tc = tc.Eval(" \"@/examples/answer.tc\" use ")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("  \"@/examples/anwer.tc\" use: %v", res)
	}
}

func TestStdlib19(t *testing.T) {
	tc := Init()
  tc = tc.Eval("use[\"@/examples/answer.tc\"]")
	if tc.Errors() != 0 {
		t.Fatalf(tc.Error())
	}
	res := tc.GetAsString()
	if res != "42" {
		t.Fatalf("use[\"@/examples/answer.tc\"] failed: %v", res)
	}
}

func TestStdlib20(t *testing.T) {
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
