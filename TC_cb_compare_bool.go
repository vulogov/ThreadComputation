package ThreadComputation

import (

)

func TCCompareBoolString(name string, x interface{}, y interface{}) interface{} {
  cfun := GetStringConverter()
  if cfun == nil {
    return nil
  }
  res := cfun(y, Bool)
  if res == nil {
    return nil
  }
  cmpfun := GetCompareCallback(x, res)
  if cmpfun == nil {
    return nil
  }
  return cmpfun(name, x, res)
}

func TCCompareStringBool(name string, x interface{}, y interface{}) interface{} {
  cfun := GetStringConverter()
  if cfun == nil {
    return nil
  }
  res := cfun(x, Bool)
  if res == nil {
    return nil
  }
  cmpfun := GetCompareCallback(res, y)
  if cmpfun == nil {
    return nil
  }
  return cmpfun(name, res, y)
}

func TCCompareBoolBool(name string, x interface{}, y interface{}) interface{} {
  switch name {
  case "=":
    return x.(bool) == y.(bool)
  case "!=":
    return x.(bool) != y.(bool)
  }
  return nil
}



func init() {
  RegisterCompareCallback(Bool, Bool, TCCompareBoolBool)
  RegisterCompareCallback(Bool, String, TCCompareBoolString)
  RegisterCompareCallback(String, Bool, TCCompareStringBool)
}
