package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func tcGetRowFromDataframe(res *TCIterator, i int64) interface{} {
  if i < 0 || i >= int64(res.Gen.(*TCData).D.NRows()) {
    res.Last = nil
    return nil
  }
  log.Debugf("dataframe-iterator: getting value at %v", i)
  e := res.Gen.(*TCData).D.Row(int(i), false)
  if e == nil {
    res.Last = nil
    return nil
  }
  out := MakeDict()
  for k, v := range e {
    out.D.Set(k,v)
  }
  res.Last = out
  return out
}

func DataGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case *TCData:
    log.Debug("iterator: create iterator for Data")
    res := new(TCIterator)
    res.Type = Data
    res.Gen  = e
    res.Last = nil
    res.Set("current", int64(0))
    return res
  }
  return nil
}

func DataNext(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  log.Debugf("%v(s) element from dataframe", i)
  out := tcGetRowFromDataframe(res, i)
  i = i + 1
  res.Set("current", i)
  return out
}

func DataPrev(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  log.Debugf("%v(s) element from list", i)
  out := tcGetRowFromDataframe(res, i)
  i = i - 1
  res.Set("current", i)
  return out
}


func init() {
  RegisterGenCallback(Data, DataGenerator)
  RegisterNextCallback(Data, DataNext)
  RegisterPrevCallback(Data, DataPrev)
}
