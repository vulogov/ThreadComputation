package ThreadComputation

import (

)



func IntGenerator(v interface{}) *TCIterator {
  switch v.(type) {
  case int64:
    res := new(TCIterator)
    res.Type = Int
    res.Gen = new(TCRange)
    res.Gen.(*TCRange).X = float64(0)
    res.Gen.(*TCRange).Y = float64(v.(int64))
    res.Gen.(*TCRange).Step = 1.0
    res.Last = nil
    return res
  }
  return nil
}

func IntNext(res *TCIterator) interface{} {
  if res.Last == nil {
    res.Last = res.Gen.(*TCRange).X
    return int64(res.Last.(float64))
  }
  res.Last = res.Last.(float64) + res.Gen.(*TCRange).Step
  if res.Last.(float64) > res.Gen.(*TCRange).Y {
    return nil
  }
  return int64(res.Last.(float64))
}

func IntPrev(res *TCIterator) interface{} {
  if res.Last == nil {
    res.Last = res.Gen.(*TCRange).Y
    return int64(res.Last.(float64))
  }
  res.Last = res.Last.(float64) - res.Gen.(*TCRange).Step
  if res.Last.(float64) < res.Gen.(*TCRange).X {
    return nil
  }
  return int64(res.Last.(float64))
}


func init() {
  RegisterGenCallback(Int, IntGenerator)
  RegisterNextCallback(Int, IntNext)
  RegisterPrevCallback(Int, IntPrev)
}
