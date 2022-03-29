package ThreadComputation

import (

)



func RangeGenerator(v interface{}) *TCIterator {
  switch v.(type) {
  case *TCRange:
    res := new(TCIterator)
    res.Type = Range
    res.Gen = v.(*TCRange)
    res.Last = nil
    return res
  }
  return nil
}

func init() {
  RegisterGenCallback(Range, RangeGenerator)
  RegisterNextCallback(Range, FloatNext)
  RegisterPrevCallback(Range, FloatPrev)
}
