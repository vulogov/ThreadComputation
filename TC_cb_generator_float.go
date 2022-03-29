package ThreadComputation

import (

)



func FloatGenerator(v interface{}) *TCIterator {
  switch v.(type) {
  case float64:
    res := new(TCIterator)
    res.Type = Float
    res.Gen = new(TCRange)
    res.Gen.(*TCRange).X = float64(0)
    res.Gen.(*TCRange).Y = v.(float64)
    res.Gen.(*TCRange).Step = 1.0
    res.Last = nil
    return res
  }
  return nil
}

func FloatNext(res *TCIterator) interface{} {
  if res.Last == nil {
    res.Last = res.Gen.(*TCRange).X
    return res.Last.(float64)
  }
  res.Last = res.Last.(float64) + res.Gen.(*TCRange).Step
  if res.Last.(float64) > res.Gen.(*TCRange).Y {
    return nil
  }
  return res.Last.(float64)
}

func FloatPrev(res *TCIterator) interface{} {
  if res.Last == nil {
    res.Last = res.Gen.(*TCRange).Y
    return res.Last.(float64)
  }
  res.Last = res.Last.(float64) - res.Gen.(*TCRange).Step
  if res.Last.(float64) < res.Gen.(*TCRange).X {
    return nil
  }
  return res.Last.(float64)
}


func init() {
  RegisterGenCallback(Float, FloatGenerator)
  RegisterNextCallback(Float, FloatNext)
  RegisterPrevCallback(Float, FloatPrev)
}
