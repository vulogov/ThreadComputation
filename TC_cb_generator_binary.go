package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func tcGetNFromBinary(res *TCIterator, i int64, s int64) interface{} {
  if i < 0 {
    return nil
  }
  log.Debugf("loop: Iterating over binary: %v %v len=%v", i, i+s, res.Gen.(*TCBinary).Len())
  if int(i + s) > res.Gen.(*TCBinary).Len() {
    return res.Gen.(*TCBinary).D[i:]
  } else {
    return res.Gen.(*TCBinary).D[i:i+s]
  }
  return nil
}

func BinaryGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case *TCBinary:
    log.Debug("iterator: create iterator for Binary")
    res := new(TCIterator)
    res.Type = Binary
    res.Gen  = e
    res.Last = nil
    res.Set("current", int64(0))
    res.Set("buffer", int64(1))
    res.Set("cross", false)
    log.Debugf("iterator: len=%v", e.Len())
    return res
  }
  return nil
}

func BinNext(res *TCIterator) interface{} {
  c := res.Get("cross").(bool)
  if c {
    return nil
  }
  i := res.Get("current").(int64)
  s := res.Get("buffer").(int64)
  if int(i) >= res.Gen.(*TCBinary).Len() {
    res.Set("cross", true)
    return nil
  }
  out := tcGetNFromBinary(res, i, s)
  i = i + s
  res.Set("current", i)
  return MakeBinary(out)
}

func BinPrev(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  s := res.Get("buffer").(int64)
  if i-s < 0 {
    i = 0
    res.Set("cross", true)
  }
  out := tcGetNFromBinary(res, i, s)
  i = i - s
  res.Set("current", i)
  return MakeBinary(out)
}


func init() {
  RegisterGenCallback(Binary, BinaryGenerator)
  RegisterNextCallback(Binary, BinNext)
  RegisterPrevCallback(Binary, BinPrev)
}
