package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

type TCBreak struct {
}

func LoopCode(l *TCExecListener, name string, code string) interface{} {
  var e interface{}
  var iter *TCIterator
  e = nil
  log.Debugf("loop %v: %v", name, code)
  if l.TC.EvAttrs.Len() > 0 {
    if l.TC.EvAttrs.Front().(*deque.Deque).Len() > 0 {
      e = l.TC.EvAttrs.Front().(*deque.Deque).Front()
    }
  }
  if l.TC.Ready() && e == nil {
    e = l.TC.Get()
  } else {
    l.TC.SetError("Context is too shallow for loop[]")
    return nil
  }
  if e == nil {
    l.TC.SetError("Context is not defined for loop[]")
    return nil
  }
  log.Debugf("Attempting to make a loop over %T", e)
  l.TC.AddContext(nil)
  switch e.(type) {
  case *TCIterator:
    iter = e.(*TCIterator)
  case *TCRange:
    r := e.(*TCRange)
    iter = l.TC.Iterator(r)
    if iter == nil {
      l.TC.DelContext()
      l.TC.SetError("Can not initiate range iterator for %T", e)
      return nil
    }
  default:
    igfun := GetGenCallback(e)
    if igfun == nil {
      l.TC.DelContext()
      l.TC.SetError("Can not initiate iterator for %T", e)
      return nil
    }
    iter = igfun(e)
    if iter == nil {
      l.TC.DelContext()
      l.TC.SetError("Can not create iterator for %T", e)
      return nil
    }  
  }
  iter.SetTC(l.TC)
  out:
  for {
    v := iter.Next()
    if v == nil {
      break out
    }
    switch v.(type) {
    case *TCBreak:
      break out
    case *TCNone:
      break out
    case *TCValue:
      switch v.(*TCValue).Value.(type) {
      case *TCNone:
        break out
      }
    }
    cfun := GetConverterCallback(e)
    if cfun == nil {
      log.Debugf("Loop element: %v", v)
    } else {
      out := cfun(v, String)
      if out == nil {
        log.Debugf("Loop element: %v", v)
      } else {
        log.Debugf("Loop element: %v", out)
      }
    }
    l.TC.AddContext(nil)
    q := new(deque.Deque)
    q.PushFront(v)
    lpfun := GetLoopPCallback(v)
    if lpfun != nil {
      q = lpfun(l, v, q)
    }
    l.TC.EvAttrs.PushFront(q)
    l.TC.Eval(code)
    l.TC.EvAttrs.PopFront()
    l.TC.DelContext()
  }
  l.TC.DelContext()
  return nil
}


func init() {
  RegisterBlockCallback("loop", LoopCode)
}
