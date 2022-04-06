package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func tcGetNFromNumbers(res *TCIterator, i int64) interface{} {
  m := int64(res.Gen.(*TCNumbers).Len())
  if i < 0 || i >= m {
    res.Last = nil
    return nil
  }
  log.Debugf("numbers-iterator: getting value at %v", i)
  return res.Gen.(*TCNumbers).N[int(i)]
}

func NumbersGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case *TCNumbers:
    log.Debug("iterator: create iterator for Numbers")
    res := new(TCIterator)
    res.Type = Numbers
    res.Gen  = e
    res.Last = nil
    res.Set("current", int64(0))
    return res
  }
  return nil
}

func NumbersNext(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  out := tcGetNFromNumbers(res, i)
  i = i + 1
  res.Set("current", i)
  return out
}

func NumbersPrev(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  out := tcGetNFromNumbers(res, i)
  i = i - 1
  res.Set("current", i)
  return out
}


func init() {
  RegisterGenCallback(Numbers, NumbersGenerator)
  RegisterNextCallback(Numbers, NumbersNext)
  RegisterPrevCallback(Numbers, NumbersPrev)
}
