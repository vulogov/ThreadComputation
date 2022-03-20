package ThreadComputation

import (
  "fmt"
  "strings"
)




func TCStringStringMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case string:
    switch y.(type) {
    case string:
      switch name {
      case "+":
        return x.(string) + y.(string)
      }
    }
  }
  return nil
}

func TCSimpleStringMath(name string, x interface{}, y interface{}) interface{} {
  switch y.(type) {
  case string:
    switch x.(type) {
    case float64:
      switch name {
      case "+":
        return fmt.Sprintf("%v%v", x.(float64), y.(string))
      case "*":
        return strings.Repeat(y.(string), int(x.(float64)))
      }
    case int64:
      switch name {
      case "+":
        return fmt.Sprintf("%v%v", x.(int64), y.(string))
      case "*":
        return strings.Repeat(y.(string), int(x.(int64)))
      }
    case bool:
      switch name {
      case "+":
        return fmt.Sprintf("%v%v", x.(bool), y.(string))
      }
    }
  }
  return nil
}

func TCStringSimpleMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case string:
    switch y.(type) {
    case float64:
      switch name {
      case "+":
        return fmt.Sprintf("%v%v", x.(string), y.(float64))
      case "*":
        return strings.Repeat(x.(string), int(y.(float64)))
      }
    case int64:
      switch name {
      case "+":
        return fmt.Sprintf("%v%v", x.(string), y.(int64))
      case "*":
        return strings.Repeat(x.(string), int(y.(int64)))
      }
    case bool:
      switch name {
      case "+":
        return fmt.Sprintf("%v%v", x.(string), y.(bool))
      }
    }
  }
  return nil
}


func init() {
  RegisterMathCallback(String, String, TCStringStringMath)
  RegisterMathCallback(String, Int, TCStringSimpleMath)
  RegisterMathCallback(String, Float, TCStringSimpleMath)
  RegisterMathCallback(String, Bool, TCStringSimpleMath)
  RegisterMathCallback(Int, String, TCSimpleStringMath)
  RegisterMathCallback(Float, String, TCSimpleStringMath)
  RegisterMathCallback(Bool, String, TCSimpleStringMath)
}
