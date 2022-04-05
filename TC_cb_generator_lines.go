package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func tcGetNFromLines(res *TCIterator, i int64) interface{} {
  if i < 0 || i >= int64(res.Gen.(*TCLines).Len()) {
    res.Last = nil
    return nil
  }
  log.Debugf("lines-iterator: getting value at %v", i)
  e := res.Gen.(*TCLines).Q.At(int(i))
  if e == nil {
    res.Last = nil
    return nil
  }
  res.Last = e
  return e
}

func LinesGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case *TCLines:
    log.Debug("iterator: create iterator for Lines")
    res := new(TCIterator)
    res.Type = Lines
    res.Gen  = e
    res.Last = nil
    res.Set("current", int64(0))
    return res
  }
  return nil
}

func LinesNext(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  log.Debugf("%v(s) element from lines", i)
  out := tcGetNFromLines(res, i)
  i = i + 1
  res.Set("current", i)
  return out
}

func LinesPrev(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  log.Debugf("%v(s) element from lines", i)
  out := tcGetNFromLines(res, i)
  i = i - 1
  res.Set("current", i)
  return out
}


func init() {
  RegisterGenCallback(Lines, LinesGenerator)
  RegisterNextCallback(Lines, LinesNext)
  RegisterPrevCallback(Lines, LinesPrev)
}
