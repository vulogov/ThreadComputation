package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func TCConversionFunction(l *TCExecListener, t int, e interface{}) (interface{}, error) {
  log.Debugf("convert[] from %v to %v", TypeToStr(e), TypeToStr(t))
  switch t {
  case List:
    lfun := GetToListCallback(e)
    if lfun == nil {
      return nil, l.TC.MakeError("conversion function to list not found")
    }
    res := lfun(e)
    if res == nil {
      return nil, l.TC.MakeError("conversion function to list returned nil")
    }
    return res, nil
  }
  fun := GetConverterCallback(e)
  if fun == nil {
    return nil, l.TC.MakeError("conversion function not found")
  }
  res := fun(e, t)
  if res == nil {
    return nil, l.TC.MakeError("conversion returned nil")
  }
  return res, nil
}

func TCConvertFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return nil, l.TC.MakeError("stack is too shallow for conversion")
  }
  rt := q.PopFront()
  t := TCType(rt)
  for q.Len() > 0 {
    e := q.PopFront()
    res, err := TCConversionFunction(l, t, e)
    if err != nil {
      return nil, err
    }
    ReturnFromFunction(l, "convert", res)
  }
  return nil, nil
}

func init() {
  SetFunction("convert", TCConvertFunction)
}
