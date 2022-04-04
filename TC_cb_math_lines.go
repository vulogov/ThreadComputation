package ThreadComputation

import (
)


func TCLinesMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCLines:
    switch y.(type) {
    case int64, float64, bool, string, *TCValue:
      switch name {
      case "+":
        x.(*TCLines).Add(y)
      }
    case *TCList:
      switch name {
      case "+":
        for i := 0; i < y.(*TCList).Len(); i++ {
          e := y.(*TCList).Q.At(i)
          x.(*TCLines).Add(e)
        }
      }
    case *TCLines:
      switch name {
      case "+":
        for i := 0; i < y.(*TCLines).Len(); i++ {
          e := y.(*TCLines).Q.At(i)
          x.(*TCLines).Add(e)
        }
      }
    }
  }
  return x
}

func init() {
  RegisterMathCallback(Lines, Int, TCLinesMath)
  RegisterMathCallback(Lines, Bool, TCLinesMath)
  RegisterMathCallback(Lines, Float, TCLinesMath)
  RegisterMathCallback(Lines, String, TCLinesMath)
  RegisterMathCallback(Lines, Value, TCLinesMath)
  RegisterMathCallback(Lines, Lines, TCLinesMath)
  RegisterMathCallback(Lines, List, TCLinesMath)
}
