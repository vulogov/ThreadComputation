package ThreadComputation

import (

)



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
}
