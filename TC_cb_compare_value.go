package ThreadComputation

import (

)



func TCCompareValueSomething(name string, x interface{}, y interface{}) interface{} {
  rx := x.(*TCValue).Value
  ry := y
  fun := GetCompareCallback(rx, ry)
  if fun == nil {
    return nil
  }
  res := fun(name, rx, ry)
  if res == nil {
    return nil
  }
  p := (x.(*TCValue).P + 100.0)/2
  return MakeValue(res.(bool), p, 0)
}

func TCCompareSomethingValue(name string, x interface{}, y interface{}) interface{} {
  ry := y.(*TCValue).Value
  rx := x
  fun := GetCompareCallback(rx, ry)
  if fun == nil {
    return nil
  }
  res := fun(name, rx, ry)
  if res == nil {
    return nil
  }
  p := (y.(*TCValue).P + 100.0)/2
  return MakeValue(res.(bool), p, 0)
}

func TCCompareValueValue(name string, x interface{}, y interface{}) interface{} {
  rx := x.(*TCValue).Value
  ry := y.(*TCValue).Value
  fun := GetCompareCallback(rx, ry)
  if fun == nil {
    return nil
  }
  res := fun(name, rx, ry)
  if res == nil {
    return nil
  }
  p := (x.(*TCValue).P + y.(*TCValue).P)/2
  return MakeValue(res.(bool), p, 0)
}

func init() {
  RegisterCompareCallback(Value, Any, TCCompareValueSomething)
  RegisterCompareCallback(Any, Value, TCCompareSomethingValue)
  RegisterCompareCallback(Value, Value, TCCompareValueValue)
}
