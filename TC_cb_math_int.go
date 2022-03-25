package ThreadComputation

import (
)




func TCIntMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case int64:
    switch y.(type) {
    case int64:
      switch name {
      case "+":
        return x.(int64) + y.(int64)
      case "-":
        return x.(int64) - y.(int64)
      case "*":
        return x.(int64) * y.(int64)
      case "/":
        if y.(int64) == 0 {
          return MakeError("Division to Zero")
        }
        return x.(int64) / y.(int64)
      }
    }
  }
  return nil
}

func init() {
  RegisterMathCallback(Int, Int, TCIntMath)
}
