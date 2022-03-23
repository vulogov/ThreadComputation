package ThreadComputation

import (
)


func listMathFun(lst *TCList, name string, y interface{}) *TCList {
  for i := 0; i < lst.Q.Len(); i++ {
    x := lst.Q.At(i)
    fun := GetMathCallback(x, y)
    if fun == nil {
      continue
    }
    e1 := fun(name, x, y)
    lst.Q.Set(i, e1)
  }
  return lst
}

func TCListMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCList:
    switch y.(type) {
    case int64:
      switch name {
      case "+":
        return listMathFun(x.(*TCList), name, y)
      case "-":
        return listMathFun(x.(*TCList), name, y)
      case "*":
        return listMathFun(x.(*TCList), name, y)
      case "/":
        return listMathFun(x.(*TCList), name, y)
      }
    }
  }
  return nil
}

func init() {
  RegisterMathCallback(List, Int, TCListMath)
  RegisterMathCallback(List, Bool, TCListMath)
  RegisterMathCallback(List, Float, TCListMath)
  RegisterMathCallback(List, String, TCListMath)
}
