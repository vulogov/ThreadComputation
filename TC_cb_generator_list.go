package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func tcGetNFromList(res *TCIterator, i int64) interface{} {
  if i < 0 || i >= int64(res.Gen.(*TCList).Len()) {
    res.Last = nil
    return nil
  }
  log.Debugf("list-iterator: getting value at %v", i)
  e := res.Gen.(*TCList).Q.At(int(i))
  if e == nil {
    res.Last = nil
    return nil
  }
  res.Last = e
  return e
}

func ListGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case *TCList:
    log.Debug("iterator: create iterator for List")
    res := new(TCIterator)
    res.Type = List
    res.Gen  = e
    res.Last = nil
    res.Set("current", int64(0))
    return res
  }
  return nil
}

func ListNext(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  out := tcGetNFromList(res, i)
  i = i + 1
  res.Set("current", i)
  return out
}

func ListPrev(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  out := tcGetNFromList(res, i)
  i = i - 1
  res.Set("current", i)
  return out
}


func init() {
  RegisterGenCallback(List, ListGenerator)
  RegisterNextCallback(List, ListNext)
  RegisterPrevCallback(List, ListPrev)
}
