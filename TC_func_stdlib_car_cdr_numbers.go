package ThreadComputation

import (

)

func TCNumbersCar(v interface{}) interface{} {
  if v.(*TCNumbers).Len() == 0 {
    return nil
  }
  return v.(*TCNumbers).N[0]
}

func TCNumbersCdr(v interface{}) interface{} {
  if v.(*TCNumbers).Len() == 0 {
    return nil
  }
  res := new(TCNumbers)
  for x := 1; x < v.(*TCNumbers).Len(); x++ {
    res.Add(v.(*TCNumbers).N[x])
  }
  return res
}

func init() {
  RegisterFunctionCallback("car", Numbers, TCNumbersCar)
  RegisterFunctionCallback("cdr", Numbers, TCNumbersCdr)
}
