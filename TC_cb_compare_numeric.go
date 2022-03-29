package ThreadComputation

import (

)

func TCCompareNumericString(name string, x interface{}, y interface{}) interface{} {
  cfun := GetStringConverter()
  if cfun == nil {
    return nil
  }
  res := cfun(y, Float)
  if res == nil {
    return nil
  }
  cmpfun := GetCompareCallback(x, res)
  if cmpfun == nil {
    return nil
  }
  return cmpfun(name, x, res)
}

func TCCompareStringNumeric(name string, x interface{}, y interface{}) interface{} {
  cfun := GetStringConverter()
  if cfun == nil {
    return nil
  }
  res := cfun(x, Float)
  if res == nil {
    return nil
  }
  cmpfun := GetCompareCallback(res, y)
  if cmpfun == nil {
    return nil
  }
  return cmpfun(name, res, y)
}

func TCCompareIntInt(name string, x interface{}, y interface{}) interface{} {
  switch name {
  case "=":
    return x.(int64) == y.(int64)
  case ">":
    return x.(int64) > y.(int64)
  case ">=":
    return x.(int64) >= y.(int64)
  case "<":
    return x.(int64) < y.(int64)
  case "<=":
    return x.(int64) <= y.(int64)
  case "!=":
    return x.(int64) != y.(int64)
  }
  return nil
}

func TCCompareIntFloat(name string, x interface{}, y interface{}) interface{} {
  switch name {
  case "=":
    return float64(x.(int64)) == y.(float64)
  case ">":
    return float64(x.(int64)) > y.(float64)
  case ">=":
    return float64(x.(int64)) >= y.(float64)
  case "<":
    return float64(x.(int64)) < y.(float64)
  case "<=":
    return float64(x.(int64)) <= y.(float64)
  case "!=":
    return float64(x.(int64)) != y.(float64)
  }
  return nil
}

func TCCompareFloatInt(name string, x interface{}, y interface{}) interface{} {
  switch name {
  case "=":
    return x.(float64) == float64(y.(int64))
  case ">":
    return x.(float64) > float64(y.(int64))
  case ">=":
    return x.(float64) >= float64(y.(int64))
  case "<":
    return x.(float64) < float64(y.(int64))
  case "<=":
    return x.(float64) <= float64(y.(int64))
  case "!=":
    return x.(float64) != float64(y.(int64))
  }
  return nil
}

func TCCompareFloatFloat(name string, x interface{}, y interface{}) interface{} {
  switch name {
  case "=":
    return x.(float64) == y.(float64)
  case ">":
    return x.(float64) > y.(float64)
  case ">=":
    return x.(float64) >= y.(float64)
  case "<":
    return x.(float64) < y.(float64)
  case "<=":
    return x.(float64) <= y.(float64)
  case "!=":
    return x.(float64) != y.(float64)
  }
  return nil
}

func init() {
  RegisterCompareCallback(Int, Int, TCCompareIntInt)
  RegisterCompareCallback(Int, Float, TCCompareIntFloat)
  RegisterCompareCallback(Float, Int, TCCompareFloatInt)
  RegisterCompareCallback(Float, Float, TCCompareFloatFloat)
  RegisterCompareCallback(Float, String, TCCompareNumericString)
  RegisterCompareCallback(Int, String, TCCompareNumericString)
  RegisterCompareCallback(String, Float, TCCompareStringNumeric)
  RegisterCompareCallback(String, Int, TCCompareStringNumeric)
}
