package ThreadComputation

import (
  "fmt"
)

type TCToStringFun func(interface{}) string


func TCConvertAnyToString(x interface{}) string {
  switch x.(type) {
  case string:
    return x.(string)
  }
  fun := GetConverterCallback(x)
  if fun == nil {
    return fmt.Sprintf("%v", x)
  }
  res := fun(x, String)
  if res == nil {
    return fmt.Sprintf("%v", x)
  }
  return res.(string)
}

func RegisterToStringCallback(x_type int, fun TCToStringFun) {
  fname := fmt.Sprintf("tostring.%v", x_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func GetToStringCallback(x interface{}) TCToStringFun {
  var fn string

  fn = fmt.Sprintf("tostring.%v", TCType(x))
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCToStringFun)
  }

  fn = fmt.Sprintf("tostring.%v", Any)
  fun, ok = Callbacks.Load(fn)
  if ok {
    return fun.(TCToStringFun)
  }

  return nil
}

func init() {
  RegisterToStringCallback(Any, TCConvertAnyToString)
}
