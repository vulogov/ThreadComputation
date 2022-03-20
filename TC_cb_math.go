package ThreadComputation

import (
  "fmt"
)



type TCMathFun func(string, interface{}, interface{}) interface{}


func GetMathCallback(x interface{}, y interface{}) TCMathFun {
  x_type := TCType(x)
  y_type := TCType(y)
  fname := fmt.Sprintf("math.%v.%v", x_type, y_type)
  if fun, ok := Callbacks.Load(fname); ok {
    return fun.(TCMathFun)
  }
  fname = fmt.Sprintf("math.%v.%v", Any, y_type)
  if fun, ok := Callbacks.Load(fname); ok {
    return fun.(TCMathFun)
  }
  fname = fmt.Sprintf("math.%v.%v", x_type, Any)
  if fun, ok := Callbacks.Load(fname); ok {
    return fun.(TCMathFun)
  }
  fname = fmt.Sprintf("math.%v.%v", Any, Any)
  if fun, ok := Callbacks.Load(fname); ok {
    return fun.(TCMathFun)
  }
  return nil
}

func RegisterMathCallback(x_type int, y_type int, fun TCMathFun) {
  fname := fmt.Sprintf("math.%v.%v", x_type, y_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
