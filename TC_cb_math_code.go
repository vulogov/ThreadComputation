package ThreadComputation

import (

)

func tcMergeCode(x *TCCode, y *TCCode) *TCCode {
  for i := 0; i < y.Code.Len(); i++ {
    x.Add(y.Code.At(i).(string))
  }
  x.Userfuncs = x.Userfuncs.Union(y.Userfuncs)
  return x
}

func TCCodeMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCCode:
    switch y.(type) {
    case int64, float64, bool, string:
      switch name {
      case "+":
        x.(*TCCode).Add(y)
        return x
      }
    case *TCCode:
      switch name {
      case "+":
        return tcMergeCode(x.(*TCCode), y.(*TCCode))
      }
    }
  }
  return nil
}

func init() {
  RegisterMathCallback(Code, Int, TCCodeMath)
  RegisterMathCallback(Code, Bool, TCCodeMath)
  RegisterMathCallback(Code, Float, TCCodeMath)
  RegisterMathCallback(Code, String, TCCodeMath)
  RegisterMathCallback(Code, Code, TCCodeMath)
}
