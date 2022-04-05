package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func StringGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case string:
    log.Debug("iterator: create iterator for String")
    res := new(TCIterator)
    res.Type = String
    res.Gen  = MakeBinary(e)
    res.Last = nil
    res.Set("current", int64(0))
    res.Set("buffer", int64(len(e)))
    res.Set("cross", false)
    log.Debugf("iterator: len=%v", len(e))
    return res
  case *TCValue:
    return StringGenerator(e.Value)
  }
  return nil
}

func StringNext(res *TCIterator) interface{} {
  out := BinNext(res)
  if out == nil {
    return nil
  }
  return string(out.(*TCBinary).D)
}

func StringPrev(res *TCIterator) interface{} {
  out := BinPrev(res)
  if out == nil {
    return nil
  }
  return string(out.(*TCBinary).D)
}

func init() {
  RegisterGenCallback(String, StringGenerator)
  RegisterNextCallback(String, StringNext)
  RegisterPrevCallback(String, StringPrev)
}
