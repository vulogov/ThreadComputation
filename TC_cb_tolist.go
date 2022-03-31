package ThreadComputation

import (
  "fmt"
)

type TCToListFun func(interface{}) *TCList


func TCConvertAnyToList(x interface{}) *TCList {
  res := new(TCList)
  res.Add(x)
  return res
}

func TCConvertNumbersToList(x interface{}) *TCList {
  switch e := x.(type) {
  case *TCNumbers:
    res := new(TCList)
    for i := 0; i < len(e.N); i++ {
      res.Add(e.N[i])
    }
    return res
  }
  return nil
}

func RegisterToListCallback(x_type int, fun TCToListFun) {
  fname := fmt.Sprintf("tolist.%v", x_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func GetToListCallback(x interface{}) TCToListFun {
  var fn string

  fn = fmt.Sprintf("tolist.%v", TCType(x))
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCToListFun)
  }

  fn = fmt.Sprintf("tolist.%v", Any)
  fun, ok = Callbacks.Load(fn)
  if ok {
    return fun.(TCToListFun)
  }

  return nil
}

func init() {
  RegisterToListCallback(Any, TCConvertAnyToList)
  RegisterToListCallback(List, TCConvertNumbersToList)
}
