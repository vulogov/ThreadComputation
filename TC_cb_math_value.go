package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)



func TCValueValueMath(name string, x interface{}, y interface{}) interface{} {
  var p float32

  p = (x.(*TCValue).P + y.(*TCValue).P)/2

  fun := GetMathCallback(x.(*TCValue).Value, y.(*TCValue).Value)
  if fun == nil {
    return MakeError("No callback for Value")
  }
  res := fun(name, x.(*TCValue).Value, y.(*TCValue).Value)
  log.Debugf("Result of math operation between Value-Value is %v", res)
  if res == nil {
    return MakeError("Value arithmetic return none")
  }
  return MakeValue(res, p, 0)
}

func TCValueAnyMath(name string, x interface{}, y interface{}) interface{} {
  var p float32
  var ry interface{}
  switch x.(type) {
  case *TCValue:
    switch y.(type) {
    case *TCValue:
      p = (x.(*TCValue).P + y.(*TCValue).P)/2
      ry = y.(*TCValue).Value
    default:
      p = (x.(*TCValue).P + float32(100.0))/2
      ry = y
    }
    fun := GetMathCallback(x.(*TCValue).Value, ry)
    if fun == nil {
      return MakeError("No callback for Value")
    }
    res := fun(name, x.(*TCValue).Value, ry)
    if res == nil {
      return MakeError("Value arithmetic return none")
    }
    return MakeValue(res, p, 0)
  }
  return nil
}

func TCAnyValueMath(name string, x interface{}, y interface{}) interface{} {
  var p float32
  var rx interface{}
  switch y.(type) {
  case *TCValue:
    switch x.(type) {
    case *TCValue:
      p = (x.(*TCValue).P + y.(*TCValue).P)/2
      rx = x.(*TCValue).Value
    default:
      p = (y.(*TCValue).P + float32(100.0))/2
      rx = x
    }
    fun := GetMathCallback(rx, y.(*TCValue).Value)
    if fun == nil {
      return MakeError("No callback for Value")
    }
    res := fun(name, rx, y.(*TCValue).Value)
    if res == nil {
      return MakeError("Value arithmetic return none")
    }
    return MakeValue(res, p, 0)
  }
  return nil
}

func init() {
  RegisterMathCallback(Value, Value, TCValueAnyMath)
  RegisterMathCallback(Value, Any, TCValueAnyMath)
  RegisterMathCallback(Any, Value, TCAnyValueMath)
}
