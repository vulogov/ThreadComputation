package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func tcGetNFromBinary(res *TCIterator, i int64, s int64) interface{} {
  return nil
}

func BinaryGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case *TCBinary:
    log.Debug("iterator: create iterator for Binary")
    res := new(TCIterator)
    res.Type = List
    res.Gen  = e
    res.Last = nil
    res.Set("current", int64(0))
    res.Set("buffer", int64(1))
    return res
  }
  return nil
}

func BinNext(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  s := res.Get("buffer").(int64)
  out := tcGetNFromBinary(res, i, s)
  i = i + 1
  res.Set("current", i)
  return out
}

func BinPrev(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  s := res.Get("buffer").(int64)
  out := tcGetNFromBinary(res, i, s)
  i = i - 1
  res.Set("current", i)
  return out
}


func init() {
  RegisterGenCallback(Binary, BinaryGenerator)
  RegisterNextCallback(Binary, BinNext)
  RegisterPrevCallback(Binary, BinPrev)
}
