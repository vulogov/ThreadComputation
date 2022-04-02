package ThreadComputation

import (

)

func TCBinaryMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCBinary:
    switch y.(type) {
    case string:
      switch name {
      case "+":
        x.(*TCBinary).Add(y)
        return x
      }
    case *TCBinary:
      switch name {
      case "+":
        return x.(*TCBinary).Add(y)
      }
    }
  }
  return nil
}

func init() {
  RegisterMathCallback(Binary, String, TCBinaryMath)
  RegisterMathCallback(Binary, Binary, TCBinaryMath)
}
