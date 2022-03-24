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

func listsMathFun(lst *TCList, name string, lst2 *TCList) *TCList {
  if lst.Q.Len() != lst2.Q.Len() {
    return nil
  }
  for i := 0; i < lst.Q.Len(); i++ {
    x := lst.Q.At(i)
    y := lst2.Q.At(i)
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
    case int64, float64, bool, string:
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
    case *TCList:
      switch name {
      case "+":
        return listsMathFun(x.(*TCList), name, y.(*TCList))
      case "-":
        return listsMathFun(x.(*TCList), name, y.(*TCList))
      case "*":
        return listsMathFun(x.(*TCList), name, y.(*TCList))
      case "/":
        return listsMathFun(x.(*TCList), name, y.(*TCList))
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
  RegisterMathCallback(List, List, TCListMath)
}
