package ThreadComputation

import (
)




func TCFloatFloatMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case float64:
    switch y.(type) {
    case float64:
      switch name {
      case "+":
        return x.(float64) + y.(float64)
      case "-":
        return x.(float64) - y.(float64)
      case "*":
        return x.(float64) * y.(float64)
      case "/":
        return x.(float64) / y.(float64)
      }
    }
  }
  return nil
}

func TCIntFloatMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case int64:
    switch y.(type) {
    case float64:
      switch name {
      case "+":
        return float64(x.(int64)) + y.(float64)
      case "-":
        return float64(x.(int64)) - y.(float64)
      case "*":
        return float64(x.(int64)) * y.(float64)
      case "/":
        return float64(x.(int64)) / y.(float64)
      }
    }
  }
  return nil
}

func TCFloatIntMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case float64:
    switch y.(type) {
    case int64:
      switch name {
      case "+":
        return x.(float64) + float64(y.(int64))
      case "-":
        return x.(float64) - float64(y.(int64))
      case "*":
        return x.(float64) * float64(y.(int64))
      case "/":
        return x.(float64) / float64(y.(int64))
      }
    }
  }
  return nil
}

func init() {
  RegisterMathCallback(Float, Float, TCFloatFloatMath)
  RegisterMathCallback(Float, Int, TCFloatIntMath)
  RegisterMathCallback(Int, Float, TCIntFloatMath)
}
