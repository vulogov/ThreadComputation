package ThreadComputation

import (
  "fmt"
)

type TCApplyFun func(interface{}, interface{}) interface{}


func TCMergeValuesToList(x interface{}, y interface{}) interface{} {
  res := &TCList{}
  res.Add(x)
  res.Add(y)
  return res
}

func TCMergeValueToListX(x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCList:
    x.(*TCList).Add(y)
  }
  return x
}

func TCMergeValueToListY(x interface{}, y interface{}) interface{} {
  switch y.(type) {
  case *TCList:
    y.(*TCList).Add(x)
  }
  return y
}

func TCMergeListsX(x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCList:
    switch y.(type) {
    case *TCList:
      for i := 0; i < x.(*TCList).Len(); i++ {
        y.(*TCList).Add(x.(*TCList).Q.At(i))
      }
    }
  }
  return y
}

func TCMergeListsY(x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCList:
    switch y.(type) {
    case *TCList:
      for i := 0; i < y.(*TCList).Len(); i++ {
        x.(*TCList).Add(y.(*TCList).Q.At(i))
      }
    }
  }
  return x
}

func RegisterApplyCallback(x_type int,y_type int, apply_fun TCApplyFun) {
  fname := fmt.Sprintf("apply.%v.%v", x_type, y_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, apply_fun)
}

func GetApplyCallback(x interface{}, y interface{}) TCApplyFun {
  var fn string

  fn = fmt.Sprintf("apply.%v.%v", TCType(x), TCType(y))

  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCApplyFun)
  }
  switch x.(type) {
  case int64,float64,string,bool:
    fn = fmt.Sprintf("apply.%v.%v", Simple, TCType(y))
    fun, ok := Callbacks.Load(fn)
    if ok {
      return fun.(TCApplyFun)
    }
  }
  switch y.(type) {
  case int64,float64,string,bool:
    fn = fmt.Sprintf("apply.%v.%v", TCType(x), Simple)
    fun, ok := Callbacks.Load(fn)
    if ok {
      return fun.(TCApplyFun)
    }
  }
  switch x.(type) {
  case int64,float64,string,bool:
    switch y.(type) {
    case int64,float64,string,bool:
      fn = fmt.Sprintf("apply.%v.%v", Simple, Simple)
      fun, ok := Callbacks.Load(fn)
      if ok {
        return fun.(TCApplyFun)
      }
    }
  }
  return nil
}

func init() {
  RegisterApplyCallback(Simple, Simple, TCMergeValuesToList)
  RegisterApplyCallback(Simple, List, TCMergeValueToListY)
  RegisterApplyCallback(List, Simple, TCMergeValueToListX)
  RegisterApplyCallback(List, List, TCMergeListsY)
}
