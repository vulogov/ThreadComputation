package ThreadComputation

import (

)

func TCListCar(v interface{}) interface{} {
  if v.(*TCList).Q.Len() == 0 {
    return nil
  }
  return v.(*TCList).Q.Front()
}

func TCListCdr(v interface{}) interface{} {
  if v.(*TCList).Q.Len() == 0 {
    return nil
  }
  res := new(TCList)
  for x := 1; x < v.(*TCList).Q.Len(); x++ {
    res.Add(v.(*TCList).Q.At(x))
  }
  return res
}

func init() {
  RegisterFunctionCallback("car", List, TCListCar)
  RegisterFunctionCallback("cdr", List, TCListCdr)
}
